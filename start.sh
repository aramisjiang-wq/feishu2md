#!/bin/bash

echo "=========================================="
echo "  飞书文档转换 Markdown 工具 - 启动脚本"
echo "=========================================="
echo ""

# 检查应用是否存在
if [ ! -f "./feishu2md4web" ]; then
    echo "⚠️  应用未编译，正在编译..."
    go build -o feishu2md4web ./web/*.go
    if [ $? -ne 0 ]; then
        echo "❌ 编译失败"
        exit 1
    fi
    echo "✅ 编译成功"
fi

# 检查是否安装了 ngrok
if ! command -v ngrok &> /dev/null; then
    echo "⚠️  ngrok 未安装"
    echo ""
    echo "请安装 ngrok："
    echo "  brew install ngrok"
    echo ""
    echo "或者访问：https://ngrok.com/download"
    exit 1
fi

echo "选择启动方式："
echo "1) 本地运行 (http://localhost:8080)"
echo "2) 本地运行 + ngrok (公网访问)"
echo ""
read -p "请输入选项 (1 或 2): " choice

case $choice in
    1)
        echo ""
        echo "🚀 启动本地应用..."
        echo "📱 访问地址: http://localhost:8080"
        echo ""
        ./feishu2md4web
        ;;
    2)
        echo ""
        echo "🚀 启动本地应用 + ngrok..."
        echo ""
        
        # 启动应用
        ./feishu2md4web &
        APP_PID=$!
        
        # 等待应用启动
        sleep 3
        
        # 启动 ngrok
        echo "📱 ngrok 地址:"
        ngrok http 8080
        
        # 清理
        kill $APP_PID 2>/dev/null
        ;;
    *)
        echo "❌ 无效选项"
        exit 1
        ;;
esac
