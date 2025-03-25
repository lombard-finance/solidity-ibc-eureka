FROM node:18-alpine

WORKDIR /hardhat

COPY . .

RUN npm install

EXPOSE 8545

CMD ["sh", "-c", "npx hardhat --verbose node --hostname 0.0.0.0"]