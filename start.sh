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

# 检查 ngrok 是否已认证（仅用于选项 2）
check_ngrok_auth() {
    if ngrok config check &> /dev/null; then
        return 0
    else
        return 1
    fi
}

# 检查是否有参数
if [ $# -eq 0 ]; then
    echo "选择启动方式："
    echo "1) 本地运行 (http://localhost:8081)"
    echo "2) 本地运行 + ngrok (公网访问)"
    echo ""
    read -p "请输入选项 (1 或 2): " choice
else
    choice="$1"
fi

case $choice in
    1)
        echo ""
        echo "🚀 启动本地应用..."
        echo "📱 访问地址: http://localhost:8081"
        echo ""
        ./feishu2md4web
        ;;
    2)
        echo ""
        echo "🚀 启动本地应用 + ngrok..."
        echo ""
        
        # 检查ngrok认证
        if ! check_ngrok_auth; then
            echo "❌ ngrok 未认证，无法使用公网访问功能"
            echo ""
            echo "请先注册并认证 ngrok："
            echo "1. 访问 https://dashboard.ngrok.com/signup 注册账号"
            echo "2. 登录后在控制台获取认证令牌"
            echo "3. 安装认证令牌：ngrok config add-authtoken YOUR_AUTH_TOKEN"
            echo ""
            exit 1
        fi
        
        # 启动应用
        ./feishu2md4web &
        APP_PID=$!
        
        # 等待应用启动
        sleep 3
        
        # 启动 ngrok
        echo "📱 ngrok 地址:"
        ngrok http 8081
        
        # 清理
        kill $APP_PID 2>/dev/null
        ;;
    *)
        echo "❌ 无效选项"
        exit 1
        ;;
esac
