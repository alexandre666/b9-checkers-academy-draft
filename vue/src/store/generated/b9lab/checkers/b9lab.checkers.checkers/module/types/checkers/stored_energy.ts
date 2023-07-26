/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "b9lab.checkers.checkers";

export interface StoredEnergy {
  index: string;
  kwh: string;
}

const baseStoredEnergy: object = { index: "", kwh: "" };

export const StoredEnergy = {
  encode(message: StoredEnergy, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.kwh !== "") {
      writer.uint32(18).string(message.kwh);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): StoredEnergy {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStoredEnergy } as StoredEnergy;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.kwh = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StoredEnergy {
    const message = { ...baseStoredEnergy } as StoredEnergy;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.kwh !== undefined && object.kwh !== null) {
      message.kwh = String(object.kwh);
    } else {
      message.kwh = "";
    }
    return message;
  },

  toJSON(message: StoredEnergy): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.kwh !== undefined && (obj.kwh = message.kwh);
    return obj;
  },

  fromPartial(object: DeepPartial<StoredEnergy>): StoredEnergy {
    const message = { ...baseStoredEnergy } as StoredEnergy;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.kwh !== undefined && object.kwh !== null) {
      message.kwh = object.kwh;
    } else {
      message.kwh = "";
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
