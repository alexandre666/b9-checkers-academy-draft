/* eslint-disable */
import { Params } from "../checkers/params";
import { SystemInfo } from "../checkers/system_info";
import { StoredGame } from "../checkers/stored_game";
import { PlayerInfo } from "../checkers/player_info";
import { Storedenergy } from "../checkers/storedenergy";
import { StoredEnergy } from "../checkers/stored_energy";
import { Energy } from "../checkers/energy";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "b9lab.checkers.checkers";

/** GenesisState defines the checkers module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  systemInfo: SystemInfo | undefined;
  storedGameList: StoredGame[];
  playerInfoList: PlayerInfo[];
  storedenergyList: Storedenergy[];
  storedEnergyLists: StoredEnergy[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  energyList: Energy[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.systemInfo !== undefined) {
      SystemInfo.encode(message.systemInfo, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.storedGameList) {
      StoredGame.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.playerInfoList) {
      PlayerInfo.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.storedenergyList) {
      Storedenergy.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.storedEnergyLists) {
      StoredEnergy.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    for (const v of message.energyList) {
      Energy.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.storedGameList = [];
    message.playerInfoList = [];
    message.storedenergyList = [];
    message.storedEnergyLists = [];
    message.energyList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.systemInfo = SystemInfo.decode(reader, reader.uint32());
          break;
        case 3:
          message.storedGameList.push(
            StoredGame.decode(reader, reader.uint32())
          );
          break;
        case 4:
          message.playerInfoList.push(
            PlayerInfo.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.storedenergyList.push(
            Storedenergy.decode(reader, reader.uint32())
          );
          break;
        case 6:
          message.storedEnergyLists.push(
            StoredEnergy.decode(reader, reader.uint32())
          );
          break;
        case 7:
          message.energyList.push(Energy.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.storedGameList = [];
    message.playerInfoList = [];
    message.storedenergyList = [];
    message.storedEnergyLists = [];
    message.energyList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.systemInfo !== undefined && object.systemInfo !== null) {
      message.systemInfo = SystemInfo.fromJSON(object.systemInfo);
    } else {
      message.systemInfo = undefined;
    }
    if (object.storedGameList !== undefined && object.storedGameList !== null) {
      for (const e of object.storedGameList) {
        message.storedGameList.push(StoredGame.fromJSON(e));
      }
    }
    if (object.playerInfoList !== undefined && object.playerInfoList !== null) {
      for (const e of object.playerInfoList) {
        message.playerInfoList.push(PlayerInfo.fromJSON(e));
      }
    }
    if (
      object.storedenergyList !== undefined &&
      object.storedenergyList !== null
    ) {
      for (const e of object.storedenergyList) {
        message.storedenergyList.push(Storedenergy.fromJSON(e));
      }
    }
    if (
      object.storedEnergyLists !== undefined &&
      object.storedEnergyLists !== null
    ) {
      for (const e of object.storedEnergyLists) {
        message.storedEnergyLists.push(StoredEnergy.fromJSON(e));
      }
    }
    if (object.energyList !== undefined && object.energyList !== null) {
      for (const e of object.energyList) {
        message.energyList.push(Energy.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.systemInfo !== undefined &&
      (obj.systemInfo = message.systemInfo
        ? SystemInfo.toJSON(message.systemInfo)
        : undefined);
    if (message.storedGameList) {
      obj.storedGameList = message.storedGameList.map((e) =>
        e ? StoredGame.toJSON(e) : undefined
      );
    } else {
      obj.storedGameList = [];
    }
    if (message.playerInfoList) {
      obj.playerInfoList = message.playerInfoList.map((e) =>
        e ? PlayerInfo.toJSON(e) : undefined
      );
    } else {
      obj.playerInfoList = [];
    }
    if (message.storedenergyList) {
      obj.storedenergyList = message.storedenergyList.map((e) =>
        e ? Storedenergy.toJSON(e) : undefined
      );
    } else {
      obj.storedenergyList = [];
    }
    if (message.storedEnergyLists) {
      obj.storedEnergyLists = message.storedEnergyLists.map((e) =>
        e ? StoredEnergy.toJSON(e) : undefined
      );
    } else {
      obj.storedEnergyLists = [];
    }
    if (message.energyList) {
      obj.energyList = message.energyList.map((e) =>
        e ? Energy.toJSON(e) : undefined
      );
    } else {
      obj.energyList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.storedGameList = [];
    message.playerInfoList = [];
    message.storedenergyList = [];
    message.storedEnergyLists = [];
    message.energyList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.systemInfo !== undefined && object.systemInfo !== null) {
      message.systemInfo = SystemInfo.fromPartial(object.systemInfo);
    } else {
      message.systemInfo = undefined;
    }
    if (object.storedGameList !== undefined && object.storedGameList !== null) {
      for (const e of object.storedGameList) {
        message.storedGameList.push(StoredGame.fromPartial(e));
      }
    }
    if (object.playerInfoList !== undefined && object.playerInfoList !== null) {
      for (const e of object.playerInfoList) {
        message.playerInfoList.push(PlayerInfo.fromPartial(e));
      }
    }
    if (
      object.storedenergyList !== undefined &&
      object.storedenergyList !== null
    ) {
      for (const e of object.storedenergyList) {
        message.storedenergyList.push(Storedenergy.fromPartial(e));
      }
    }
    if (
      object.storedEnergyLists !== undefined &&
      object.storedEnergyLists !== null
    ) {
      for (const e of object.storedEnergyLists) {
        message.storedEnergyLists.push(StoredEnergy.fromPartial(e));
      }
    }
    if (object.energyList !== undefined && object.energyList !== null) {
      for (const e of object.energyList) {
        message.energyList.push(Energy.fromPartial(e));
      }
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
