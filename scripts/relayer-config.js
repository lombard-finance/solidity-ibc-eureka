const fs = require('fs').promises;
const { StargateClient } = require('@cosmjs/stargate');

async function main() {
  const addresses = JSON.parse(await fs.readFile('./scripts/addresses.json', 'utf8'));
  const faucet = JSON.parse(await fs.readFile('./scripts/faucet.json', 'utf8'));
  const relayerConfig = JSON.parse(await fs.readFile('./scripts/relayer-config.template.json', 'utf8'));

  relayerConfig.modules.forEach((module) => module.config.ics26_address = addresses.ics26Router);
  relayerConfig.modules.forEach((module) => {
    if (module.name === "eth_to_cosmos") {
      module.config.signer_address = faucet.address;
    }
  })

  const path = './.relayer';
  try {
    await fs.access(path, fs.constants.W_OK);
  } catch (error) {
    if (error.code === 'ENOENT') {
      await fs.mkdir(path, { recursive: true });
    } else {
      throw error;
    }
  }

  await fs.writeFile('./.relayer/relayer-config.json', JSON.stringify(relayerConfig, null, 2));
}

main().then();