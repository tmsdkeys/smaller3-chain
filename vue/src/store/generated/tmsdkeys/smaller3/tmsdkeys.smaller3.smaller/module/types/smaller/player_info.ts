/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "tmsdkeys.smaller3.smaller";

export interface PlayerInfo {
  player: string;
  winCount: number;
  hasBeenTopRank: boolean;
}

const basePlayerInfo: object = {
  player: "",
  winCount: 0,
  hasBeenTopRank: false,
};

export const PlayerInfo = {
  encode(message: PlayerInfo, writer: Writer = Writer.create()): Writer {
    if (message.player !== "") {
      writer.uint32(10).string(message.player);
    }
    if (message.winCount !== 0) {
      writer.uint32(16).uint64(message.winCount);
    }
    if (message.hasBeenTopRank === true) {
      writer.uint32(24).bool(message.hasBeenTopRank);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): PlayerInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePlayerInfo } as PlayerInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.player = reader.string();
          break;
        case 2:
          message.winCount = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.hasBeenTopRank = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PlayerInfo {
    const message = { ...basePlayerInfo } as PlayerInfo;
    if (object.player !== undefined && object.player !== null) {
      message.player = String(object.player);
    } else {
      message.player = "";
    }
    if (object.winCount !== undefined && object.winCount !== null) {
      message.winCount = Number(object.winCount);
    } else {
      message.winCount = 0;
    }
    if (object.hasBeenTopRank !== undefined && object.hasBeenTopRank !== null) {
      message.hasBeenTopRank = Boolean(object.hasBeenTopRank);
    } else {
      message.hasBeenTopRank = false;
    }
    return message;
  },

  toJSON(message: PlayerInfo): unknown {
    const obj: any = {};
    message.player !== undefined && (obj.player = message.player);
    message.winCount !== undefined && (obj.winCount = message.winCount);
    message.hasBeenTopRank !== undefined &&
      (obj.hasBeenTopRank = message.hasBeenTopRank);
    return obj;
  },

  fromPartial(object: DeepPartial<PlayerInfo>): PlayerInfo {
    const message = { ...basePlayerInfo } as PlayerInfo;
    if (object.player !== undefined && object.player !== null) {
      message.player = object.player;
    } else {
      message.player = "";
    }
    if (object.winCount !== undefined && object.winCount !== null) {
      message.winCount = object.winCount;
    } else {
      message.winCount = 0;
    }
    if (object.hasBeenTopRank !== undefined && object.hasBeenTopRank !== null) {
      message.hasBeenTopRank = object.hasBeenTopRank;
    } else {
      message.hasBeenTopRank = false;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
