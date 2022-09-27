/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "tmsdkeys.smaller3.smaller";

export interface MsgPlayGame {
  creator: string;
  numPick: string;
}

export interface MsgPlayGameResponse {
  gameIndex: string;
  win: boolean;
}

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

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  PlayGame(request: MsgPlayGame): Promise<MsgPlayGameResponse>;
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
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
