/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "tmsdkeys.smaller3.smaller";

export interface SmallerPacketData {
  noData: NoData | undefined;
  /** this line is used by starport scaffolding # ibc/packet/proto/field */
  topRankPacket: TopRankPacketData | undefined;
  /** this line is used by starport scaffolding # ibc/packet/proto/field/number */
  gameResultPacket: GameResultPacketData | undefined;
}

export interface NoData {}

/** GameResultPacketData defines a struct for the packet payload */
export interface GameResultPacketData {
  gameId: number;
}

/** GameResultPacketAck defines a struct for the packet acknowledgment */
export interface GameResultPacketAck {
  gameId: number;
}

/** TopRankPacketData defines a struct for the packet payload */
export interface TopRankPacketData {
  address: string;
  score: number;
}

/** TopRankPacketAck defines a struct for the packet acknowledgment */
export interface TopRankPacketAck {
  clientId: string;
  address: string;
  hasBeenTopRank: boolean;
}

const baseSmallerPacketData: object = {};

export const SmallerPacketData = {
  encode(message: SmallerPacketData, writer: Writer = Writer.create()): Writer {
    if (message.noData !== undefined) {
      NoData.encode(message.noData, writer.uint32(10).fork()).ldelim();
    }
    if (message.topRankPacket !== undefined) {
      TopRankPacketData.encode(
        message.topRankPacket,
        writer.uint32(26).fork()
      ).ldelim();
    }
    if (message.gameResultPacket !== undefined) {
      GameResultPacketData.encode(
        message.gameResultPacket,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): SmallerPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSmallerPacketData } as SmallerPacketData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.noData = NoData.decode(reader, reader.uint32());
          break;
        case 3:
          message.topRankPacket = TopRankPacketData.decode(
            reader,
            reader.uint32()
          );
          break;
        case 2:
          message.gameResultPacket = GameResultPacketData.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SmallerPacketData {
    const message = { ...baseSmallerPacketData } as SmallerPacketData;
    if (object.noData !== undefined && object.noData !== null) {
      message.noData = NoData.fromJSON(object.noData);
    } else {
      message.noData = undefined;
    }
    if (object.topRankPacket !== undefined && object.topRankPacket !== null) {
      message.topRankPacket = TopRankPacketData.fromJSON(object.topRankPacket);
    } else {
      message.topRankPacket = undefined;
    }
    if (
      object.gameResultPacket !== undefined &&
      object.gameResultPacket !== null
    ) {
      message.gameResultPacket = GameResultPacketData.fromJSON(
        object.gameResultPacket
      );
    } else {
      message.gameResultPacket = undefined;
    }
    return message;
  },

  toJSON(message: SmallerPacketData): unknown {
    const obj: any = {};
    message.noData !== undefined &&
      (obj.noData = message.noData ? NoData.toJSON(message.noData) : undefined);
    message.topRankPacket !== undefined &&
      (obj.topRankPacket = message.topRankPacket
        ? TopRankPacketData.toJSON(message.topRankPacket)
        : undefined);
    message.gameResultPacket !== undefined &&
      (obj.gameResultPacket = message.gameResultPacket
        ? GameResultPacketData.toJSON(message.gameResultPacket)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<SmallerPacketData>): SmallerPacketData {
    const message = { ...baseSmallerPacketData } as SmallerPacketData;
    if (object.noData !== undefined && object.noData !== null) {
      message.noData = NoData.fromPartial(object.noData);
    } else {
      message.noData = undefined;
    }
    if (object.topRankPacket !== undefined && object.topRankPacket !== null) {
      message.topRankPacket = TopRankPacketData.fromPartial(
        object.topRankPacket
      );
    } else {
      message.topRankPacket = undefined;
    }
    if (
      object.gameResultPacket !== undefined &&
      object.gameResultPacket !== null
    ) {
      message.gameResultPacket = GameResultPacketData.fromPartial(
        object.gameResultPacket
      );
    } else {
      message.gameResultPacket = undefined;
    }
    return message;
  },
};

const baseNoData: object = {};

export const NoData = {
  encode(_: NoData, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NoData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNoData } as NoData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): NoData {
    const message = { ...baseNoData } as NoData;
    return message;
  },

  toJSON(_: NoData): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<NoData>): NoData {
    const message = { ...baseNoData } as NoData;
    return message;
  },
};

const baseGameResultPacketData: object = { gameId: 0 };

export const GameResultPacketData = {
  encode(
    message: GameResultPacketData,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.gameId !== 0) {
      writer.uint32(8).uint64(message.gameId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GameResultPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGameResultPacketData } as GameResultPacketData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gameId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GameResultPacketData {
    const message = { ...baseGameResultPacketData } as GameResultPacketData;
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = Number(object.gameId);
    } else {
      message.gameId = 0;
    }
    return message;
  },

  toJSON(message: GameResultPacketData): unknown {
    const obj: any = {};
    message.gameId !== undefined && (obj.gameId = message.gameId);
    return obj;
  },

  fromPartial(object: DeepPartial<GameResultPacketData>): GameResultPacketData {
    const message = { ...baseGameResultPacketData } as GameResultPacketData;
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = object.gameId;
    } else {
      message.gameId = 0;
    }
    return message;
  },
};

const baseGameResultPacketAck: object = { gameId: 0 };

export const GameResultPacketAck = {
  encode(
    message: GameResultPacketAck,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.gameId !== 0) {
      writer.uint32(8).uint64(message.gameId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GameResultPacketAck {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGameResultPacketAck } as GameResultPacketAck;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gameId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GameResultPacketAck {
    const message = { ...baseGameResultPacketAck } as GameResultPacketAck;
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = Number(object.gameId);
    } else {
      message.gameId = 0;
    }
    return message;
  },

  toJSON(message: GameResultPacketAck): unknown {
    const obj: any = {};
    message.gameId !== undefined && (obj.gameId = message.gameId);
    return obj;
  },

  fromPartial(object: DeepPartial<GameResultPacketAck>): GameResultPacketAck {
    const message = { ...baseGameResultPacketAck } as GameResultPacketAck;
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = object.gameId;
    } else {
      message.gameId = 0;
    }
    return message;
  },
};

const baseTopRankPacketData: object = { address: "", score: 0 };

export const TopRankPacketData = {
  encode(message: TopRankPacketData, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.score !== 0) {
      writer.uint32(16).uint64(message.score);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): TopRankPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTopRankPacketData } as TopRankPacketData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.score = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TopRankPacketData {
    const message = { ...baseTopRankPacketData } as TopRankPacketData;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.score !== undefined && object.score !== null) {
      message.score = Number(object.score);
    } else {
      message.score = 0;
    }
    return message;
  },

  toJSON(message: TopRankPacketData): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.score !== undefined && (obj.score = message.score);
    return obj;
  },

  fromPartial(object: DeepPartial<TopRankPacketData>): TopRankPacketData {
    const message = { ...baseTopRankPacketData } as TopRankPacketData;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.score !== undefined && object.score !== null) {
      message.score = object.score;
    } else {
      message.score = 0;
    }
    return message;
  },
};

const baseTopRankPacketAck: object = {
  clientId: "",
  address: "",
  hasBeenTopRank: false,
};

export const TopRankPacketAck = {
  encode(message: TopRankPacketAck, writer: Writer = Writer.create()): Writer {
    if (message.clientId !== "") {
      writer.uint32(10).string(message.clientId);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.hasBeenTopRank === true) {
      writer.uint32(24).bool(message.hasBeenTopRank);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): TopRankPacketAck {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTopRankPacketAck } as TopRankPacketAck;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.clientId = reader.string();
          break;
        case 2:
          message.address = reader.string();
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

  fromJSON(object: any): TopRankPacketAck {
    const message = { ...baseTopRankPacketAck } as TopRankPacketAck;
    if (object.clientId !== undefined && object.clientId !== null) {
      message.clientId = String(object.clientId);
    } else {
      message.clientId = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.hasBeenTopRank !== undefined && object.hasBeenTopRank !== null) {
      message.hasBeenTopRank = Boolean(object.hasBeenTopRank);
    } else {
      message.hasBeenTopRank = false;
    }
    return message;
  },

  toJSON(message: TopRankPacketAck): unknown {
    const obj: any = {};
    message.clientId !== undefined && (obj.clientId = message.clientId);
    message.address !== undefined && (obj.address = message.address);
    message.hasBeenTopRank !== undefined &&
      (obj.hasBeenTopRank = message.hasBeenTopRank);
    return obj;
  },

  fromPartial(object: DeepPartial<TopRankPacketAck>): TopRankPacketAck {
    const message = { ...baseTopRankPacketAck } as TopRankPacketAck;
    if (object.clientId !== undefined && object.clientId !== null) {
      message.clientId = object.clientId;
    } else {
      message.clientId = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
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
