// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateEnergy } from "./types/checkers/tx";
import { MsgCreateGame } from "./types/checkers/tx";
import { MsgPlayMove } from "./types/checkers/tx";


const types = [
  ["/b9lab.checkers.checkers.MsgCreateEnergy", MsgCreateEnergy],
  ["/b9lab.checkers.checkers.MsgCreateGame", MsgCreateGame],
  ["/b9lab.checkers.checkers.MsgPlayMove", MsgPlayMove],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgCreateEnergy: (data: MsgCreateEnergy): EncodeObject => ({ typeUrl: "/b9lab.checkers.checkers.MsgCreateEnergy", value: MsgCreateEnergy.fromPartial( data ) }),
    msgCreateGame: (data: MsgCreateGame): EncodeObject => ({ typeUrl: "/b9lab.checkers.checkers.MsgCreateGame", value: MsgCreateGame.fromPartial( data ) }),
    msgPlayMove: (data: MsgPlayMove): EncodeObject => ({ typeUrl: "/b9lab.checkers.checkers.MsgPlayMove", value: MsgPlayMove.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
