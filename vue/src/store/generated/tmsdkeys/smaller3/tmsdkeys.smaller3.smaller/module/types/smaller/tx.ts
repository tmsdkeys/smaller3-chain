/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "tmsdkeys.smaller3.smaller";

export interface MsgPlayGame {
  creator: string;
  numPick: string;
}

export interface MsgPlayGameResponse {
  gameIndex: string;
  win: boolean;
}

export interface MsgSendGameResult {
  creator: string;
  port: string;
  channelID: string;
  timeoutTimestamp: number;
  gameId: number;
}

export interface MsgSendGameResultResponse {}

const baseMsgPlayGame: object = { creator: "", numPick: "" };

export const MsgPlayGame = {
  encode(message: MsgPlayGame, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.numPick !== "") {
      writer.uint32(18).string(message.numPick);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgPlayGame {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgPlayGame } as MsgPlayGame;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.numPick = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPlayGame {
    const message = { ...baseMsgPlayGame } as MsgPlayGame;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.numPick !== undefined && object.numPick !== null) {
      message.numPick = String(object.numPick);
    } else {
      message.numPick = "";
    }
    return message;
  },

  toJSON(message: MsgPlayGame): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.numPick !== undefined && (obj.numPick = message.numPick);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgPlayGame>): MsgPlayGame {
    const message = { ...baseMsgPlayGame } as MsgPlayGame;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.numPick !== undefined && object.numPick !== null) {
      message.numPick = object.numPick;
    } else {
      message.numPick = "";
    }
    return message;
  },
};

const baseMsgPlayGameResponse: object = { gameIndex: "", win: false };

export const MsgPlayGameResponse = {
  encode(
    message: MsgPlayGameResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.gameIndex !== "") {
      writer.uint32(10).string(message.gameIndex);
    }
    if (message.win === true) {
      writer.uint32(16).bool(message.win);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgPlayGameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgPlayGameResponse } as MsgPlayGameResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gameIndex = reader.string();
          break;
        case 2:
          message.win = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPlayGameResponse {
    const message = { ...baseMsgPlayGameResponse } as MsgPlayGameResponse;
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = String(object.gameIndex);
    } else {
      message.gameIndex = "";
    }
    if (object.win !== undefined && object.win !== null) {
      message.win = Boolean(object.win);
    } else {
      message.win = false;
    }
    return message;
  },

  toJSON(message: MsgPlayGameResponse): unknown {
    const obj: any = {};
    message.gameIndex !== undefined && (obj.gameIndex = message.gameIndex);
    message.win !== undefined && (obj.win = message.win);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgPlayGameResponse>): MsgPlayGameResponse {
    const message = { ...baseMsgPlayGameResponse } as MsgPlayGameResponse;
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = object.gameIndex;
    } else {
      message.gameIndex = "";
    }
    if (object.win !== undefined && object.win !== null) {
      message.win = object.win;
    } else {
      message.win = false;
    }
    return message;
  },
};

const baseMsgSendGameResult: object = {
  creator: "",
  port: "",
  channelID: "",
  timeoutTimestamp: 0,
  gameId: 0,
};

export const MsgSendGameResult = {
  encode(message: MsgSendGameResult, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.port !== "") {
      writer.uint32(18).string(message.port);
    }
    if (message.channelID !== "") {
      writer.uint32(26).string(message.channelID);
    }
    if (message.timeoutTimestamp !== 0) {
      writer.uint32(32).uint64(message.timeoutTimestamp);
    }
    if (message.gameId !== 0) {
      writer.uint32(40).uint64(message.gameId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSendGameResult {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSendGameResult } as MsgSendGameResult;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.port = reader.string();
          break;
        case 3:
          message.channelID = reader.string();
          break;
        case 4:
          message.timeoutTimestamp = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.gameId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendGameResult {
    const message = { ...baseMsgSendGameResult } as MsgSendGameResult;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.port !== undefined && object.port !== null) {
      message.port = String(object.port);
    } else {
      message.port = "";
    }
    if (object.channelID !== undefined && object.channelID !== null) {
      message.channelID = String(object.channelID);
    } else {
      message.channelID = "";
    }
    if (
      object.timeoutTimestamp !== undefined &&
      object.timeoutTimestamp !== null
    ) {
      message.timeoutTimestamp = Number(object.timeoutTimestamp);
    } else {
      message.timeoutTimestamp = 0;
    }
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = Number(object.gameId);
    } else {
      message.gameId = 0;
    }
    return message;
  },

  toJSON(message: MsgSendGameResult): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.port !== undefined && (obj.port = message.port);
    message.channelID !== undefined && (obj.channelID = message.channelID);
    message.timeoutTimestamp !== undefined &&
      (obj.timeoutTimestamp = message.timeoutTimestamp);
    message.gameId !== undefined && (obj.gameId = message.gameId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSendGameResult>): MsgSendGameResult {
    const message = { ...baseMsgSendGameResult } as MsgSendGameResult;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.port !== undefined && object.port !== null) {
      message.port = object.port;
    } else {
      message.port = "";
    }
    if (object.channelID !== undefined && object.channelID !== null) {
      message.channelID = object.channelID;
    } else {
      message.channelID = "";
    }
    if (
      object.timeoutTimestamp !== undefined &&
      object.timeoutTimestamp !== null
    ) {
      message.timeoutTimestamp = object.timeoutTimestamp;
    } else {
      message.timeoutTimestamp = 0;
    }
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = object.gameId;
    } else {
      message.gameId = 0;
    }
    return message;
  },
};

const baseMsgSendGameResultResponse: object = {};

export const MsgSendGameResultResponse = {
  encode(
    _: MsgSendGameResultResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSendGameResultResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSendGameResultResponse,
    } as MsgSendGameResultResponse;
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

  fromJSON(_: any): MsgSendGameResultResponse {
    const message = {
      ...baseMsgSendGameResultResponse,
    } as MsgSendGameResultResponse;
    return message;
  },

  toJSON(_: MsgSendGameResultResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSendGameResultResponse>
  ): MsgSendGameResultResponse {
    const message = {
      ...baseMsgSendGameResultResponse,
    } as MsgSendGameResultResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  PlayGame(request: MsgPlayGame): Promise<MsgPlayGameResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SendGameResult(
    request: MsgSendGameResult
  ): Promise<MsgSendGameResultResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  PlayGame(request: MsgPlayGame): Promise<MsgPlayGameResponse> {
    const data = MsgPlayGame.encode(request).finish();
    const promise = this.rpc.request(
      "tmsdkeys.smaller3.smaller.Msg",
      "PlayGame",
      data
    );
    return promise.then((data) => MsgPlayGameResponse.decode(new Reader(data)));
  }

  SendGameResult(
    request: MsgSendGameResult
  ): Promise<MsgSendGameResultResponse> {
    const data = MsgSendGameResult.encode(request).finish();
    const promise = this.rpc.request(
      "tmsdkeys.smaller3.smaller.Msg",
      "SendGameResult",
      data
    );
    return promise.then((data) =>
      MsgSendGameResultResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
