# Node.js の公式イメージを使用
FROM node:20

# 作業ディレクトリ
WORKDIR /app/mobile

# 依存関係をインストール
COPY package.json package-lock.json ./
RUN npm install

# ソースコードをコピー
COPY . .

# 開発サーバーを起動
CMD ["npx", "expo", "start"]