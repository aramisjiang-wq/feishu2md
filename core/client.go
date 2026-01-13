package core

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/chai2010/webp"
	"github.com/chyroc/lark"
	"github.com/chyroc/lark_rate_limiter"
	"github.com/disintegration/imaging"
)

type Client struct {
	larkClient *lark.Lark
}

func NewClient(appID, appSecret string) *Client {
	return &Client{
		larkClient: lark.New(
			lark.WithAppCredential(appID, appSecret),
			lark.WithTimeout(60*time.Second),
			lark.WithApiMiddleware(lark_rate_limiter.Wait(4, 4)),
		),
	}
}

func (c *Client) DownloadImage(ctx context.Context, imgToken, outDir string) (string, error) {
	resp, _, err := c.larkClient.Drive.DownloadDriveMedia(ctx, &lark.DownloadDriveMediaReq{
		FileToken: imgToken,
	})
	if err != nil {
		return imgToken, err
	}
	fileext := filepath.Ext(resp.Filename)
	filename := fmt.Sprintf("%s/%s%s", outDir, imgToken, fileext)
	err = os.MkdirAll(filepath.Dir(filename), 0o755)
	if err != nil {
		return imgToken, err
	}
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		return imgToken, err
	}
	defer file.Close()
	_, err = io.Copy(file, resp.File)
	if err != nil {
		return imgToken, err
	}
	return filename, nil
}

func (c *Client) DownloadImageRaw(ctx context.Context, imgToken, imgDir string) (string, []byte, error) {
	resp, _, err := c.larkClient.Drive.DownloadDriveMedia(ctx, &lark.DownloadDriveMediaReq{
		FileToken: imgToken,
	})
	if err != nil {
		return imgToken, nil, err
	}
	fileext := filepath.Ext(resp.Filename)
	filename := fmt.Sprintf("%s/%s%s", imgDir, imgToken, fileext)
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.File)
	return filename, buf.Bytes(), nil
}

func (c *Client) DownloadImageBase64(ctx context.Context, imgToken string) (string, error) {
	resp, _, err := c.larkClient.Drive.DownloadDriveMedia(ctx, &lark.DownloadDriveMediaReq{
		FileToken: imgToken,
	})
	if err != nil {
		return "", err
	}
	
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.File)
	
	fileext := filepath.Ext(resp.Filename)
	mimeType := "image/png"
	if fileext == ".jpg" || fileext == ".jpeg" {
		mimeType = "image/jpeg"
	} else if fileext == ".gif" {
		mimeType = "image/gif"
	} else if fileext == ".webp" {
		mimeType = "image/webp"
	} else if fileext == ".svg" {
		mimeType = "image/svg+xml"
	}
	
	processedMimeType, base64Str, err := processImageForBase64(buf.Bytes(), mimeType)
	if err != nil {
		return "", err
	}
	
	return fmt.Sprintf("data:%s;base64,%s", processedMimeType, base64Str), nil
}

func (c *Client) GetDocxContent(ctx context.Context, docToken string) (*lark.DocxDocument, []*lark.DocxBlock, error) {
	resp, _, err := c.larkClient.Drive.GetDocxDocument(ctx, &lark.GetDocxDocumentReq{
		DocumentID: docToken,
	})
	if err != nil {
		return nil, nil, err
	}
	docx := &lark.DocxDocument{
		DocumentID: resp.Document.DocumentID,
		RevisionID: resp.Document.RevisionID,
		Title:      resp.Document.Title,
	}
	var blocks []*lark.DocxBlock
	var pageToken *string
	for {
		resp2, _, err := c.larkClient.Drive.GetDocxBlockListOfDocument(ctx, &lark.GetDocxBlockListOfDocumentReq{
			DocumentID: docx.DocumentID,
			PageToken:  pageToken,
		})
		if err != nil {
			return docx, nil, err
		}
		blocks = append(blocks, resp2.Items...)
		pageToken = &resp2.PageToken
		if !resp2.HasMore {
			break
		}
	}
	return docx, blocks, nil
}

func (c *Client) GetWikiNodeInfo(ctx context.Context, token string) (*lark.GetWikiNodeRespNode, error) {
	resp, _, err := c.larkClient.Drive.GetWikiNode(ctx, &lark.GetWikiNodeReq{
		Token: token,
	})
	if err != nil {
		return nil, err
	}
	return resp.Node, nil
}

func (c *Client) GetDriveFolderFileList(ctx context.Context, pageToken *string, folderToken *string) ([]*lark.GetDriveFileListRespFile, error) {
	resp, _, err := c.larkClient.Drive.GetDriveFileList(ctx, &lark.GetDriveFileListReq{
		PageSize:    nil,
		PageToken:   pageToken,
		FolderToken: folderToken,
	})
	if err != nil {
		return nil, err
	}
	files := resp.Files
	for resp.HasMore {
		resp, _, err = c.larkClient.Drive.GetDriveFileList(ctx, &lark.GetDriveFileListReq{
			PageSize:    nil,
			PageToken:   &resp.NextPageToken,
			FolderToken: folderToken,
		})
		if err != nil {
			return nil, err
		}
		files = append(files, resp.Files...)
	}
	return files, nil
}

func (c *Client) GetWikiName(ctx context.Context, spaceID string) (string, error) {
	resp, _, err := c.larkClient.Drive.GetWikiSpace(ctx, &lark.GetWikiSpaceReq{
		SpaceID: spaceID,
	})

	if err != nil {
		return "", err
	}

	return resp.Space.Name, nil
}

func (c *Client) GetWikiNodeList(ctx context.Context, spaceID string, parentNodeToken *string) ([]*lark.GetWikiNodeListRespItem, error) {
	resp, _, err := c.larkClient.Drive.GetWikiNodeList(ctx, &lark.GetWikiNodeListReq{
		SpaceID:         spaceID,
		PageSize:        nil,
		PageToken:       nil,
		ParentNodeToken: parentNodeToken,
	})

	if err != nil {
		return nil, err
	}

	nodes := resp.Items
	previousPageToken := ""

	for resp.HasMore && previousPageToken != resp.PageToken {
		previousPageToken = resp.PageToken
		resp, _, err := c.larkClient.Drive.GetWikiNodeList(ctx, &lark.GetWikiNodeListReq{
			SpaceID:         spaceID,
			PageSize:        nil,
			PageToken:       &resp.PageToken,
			ParentNodeToken: parentNodeToken,
		})

		if err != nil {
			return nil, err
		}

		nodes = append(nodes, resp.Items...)
	}

	return nodes, nil
}

const (
	maxImageWidth  = 800
	imageQuality   = 75
	webpQuality    = 75
)

func compressImage(img image.Image) image.Image {
	bounds := img.Bounds()
	width := bounds.Dx()
	
	if width > maxImageWidth {
		return imaging.Resize(img, maxImageWidth, 0, imaging.Lanczos)
	}
	return img
}

func encodeToWebP(img image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := webp.Encode(buf, img, &webp.Options{Lossless: false, Quality: float32(webpQuality)})
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func encodeToJPEG(img image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, img, &jpeg.Options{Quality: imageQuality})
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func processImageForBase64(data []byte, mimeType string) (string, string, error) {
	originalSize := len(data)
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return "", "", fmt.Errorf("failed to decode image: %w", err)
	}

	compressed := compressImage(img)
	webpData, err := encodeToWebP(compressed)
	if err != nil {
		return "", "", fmt.Errorf("failed to encode to WebP: %w", err)
	}

	compressedSize := len(webpData)
	reduction := float64(originalSize-compressedSize) / float64(originalSize) * 100
	fmt.Printf("[图片压缩] 原始: %d bytes, 压缩后: %d bytes, 减少: %.1f%%\n", originalSize, compressedSize, reduction)

	base64Str := base64.StdEncoding.EncodeToString(webpData)
	return "image/webp", base64Str, nil
}
