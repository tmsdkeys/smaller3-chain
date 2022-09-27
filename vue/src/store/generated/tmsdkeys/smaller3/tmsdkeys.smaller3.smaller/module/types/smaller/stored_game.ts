/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "tmsdkeys.smaller3.smaller";

export interface StoredGame {
  index: string;
  player: string;
  win: boolean;
  sentToLB: boolean;
}

const baseStoredGame: object = {
  index: "",
  player: "",
  win: false,
  sentToLB: false,
};

export const StoredGame = {
  encode(message: StoredGame, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.player !== "") {
      writer.uint32(18).string(message.player);
    }
    if (message.win === true) {
      writer.uint32(24).bool(message.win);
    }
    if (message.sentToLB === true) {
      writer.uint32(32).bool(message.sentToLB);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): StoredGame {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStoredGame } as StoredGame;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.player = reader.string();
          break;
        case 3:
          message.win = reader.bool();
          break;
        case 4:
          message.sentToLB = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StoredGame {
    const message = { ...baseStoredGame } as StoredGame;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.player !== undefined && object.player !== null) {
      message.player = String(object.player);
    } else {
      message.player = "";
    }
    if (object.win !== undefined && object.win !== null) {
      message.win = Boolean(object.win);
    } else {
      message.win = false;
    }
    if (object.sentToLB !== undefined && object.sentToLB !== null) {
      message.sentToLB = Boolean(object.sentToLB);
    } else {
      message.sentToLB = false;
    }
    return message;
  },

  toJSON(message: StoredGame): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.player !== undefined && (obj.player = message.player);
    message.win !== undefined && (obj.win = message.win);
    message.sentToLB !== undefined && (obj.sentToLB = message.sentToLB);
    return obj;
  },

  fromPartial(object: DeepPartial<StoredGame>): StoredGame {
    const message = { ...baseStoredGame } as StoredGame;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.player !== undefined && object.player !== null) {
      message.player = object.player;
    } else {
      message.player = "";
    }
    if (object.win !== undefined && object.win !== null) {
      message.win = object.win;
    } else {
      message.win = false;
    }
    if (object.sentToLB !== undefined && object.sentToLB !== null) {
      message.sentToLB = object.sentToLB;
    } else {
      message.sentToLB = false;
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
