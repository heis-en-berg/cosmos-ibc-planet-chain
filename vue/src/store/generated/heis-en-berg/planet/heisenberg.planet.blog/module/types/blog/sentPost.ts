/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'heisenberg.planet.blog'

export interface SentPost {
  creator: string
  id: number
  postId: string
  title: string
  chain: string
}

const baseSentPost: object = { creator: '', id: 0, postId: '', title: '', chain: '' }

export const SentPost = {
  encode(message: SentPost, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id)
    }
    if (message.postId !== '') {
      writer.uint32(26).string(message.postId)
    }
    if (message.title !== '') {
      writer.uint32(34).string(message.title)
    }
    if (message.chain !== '') {
      writer.uint32(42).string(message.chain)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): SentPost {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseSentPost } as SentPost
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.id = longToNumber(reader.uint64() as Long)
          break
        case 3:
          message.postId = reader.string()
          break
        case 4:
          message.title = reader.string()
          break
        case 5:
          message.chain = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): SentPost {
    const message = { ...baseSentPost } as SentPost
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id)
    } else {
      message.id = 0
    }
    if (object.postId !== undefined && object.postId !== null) {
      message.postId = String(object.postId)
    } else {
      message.postId = ''
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title)
    } else {
      message.title = ''
    }
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = String(object.chain)
    } else {
      message.chain = ''
    }
    return message
  },

  toJSON(message: SentPost): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.id !== undefined && (obj.id = message.id)
    message.postId !== undefined && (obj.postId = message.postId)
    message.title !== undefined && (obj.title = message.title)
    message.chain !== undefined && (obj.chain = message.chain)
    return obj
  },

  fromPartial(object: DeepPartial<SentPost>): SentPost {
    const message = { ...baseSentPost } as SentPost
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = 0
    }
    if (object.postId !== undefined && object.postId !== null) {
      message.postId = object.postId
    } else {
      message.postId = ''
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title
    } else {
      message.title = ''
    }
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = object.chain
    } else {
      message.chain = ''
    }
    return message
  }
}

declare var self: any | undefined
declare var window: any | undefined
var globalThis: any = (() => {
  if (typeof globalThis !== 'undefined') return globalThis
  if (typeof self !== 'undefined') return self
  if (typeof window !== 'undefined') return window
  if (typeof global !== 'undefined') return global
  throw 'Unable to locate global object'
})()

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER')
  }
  return long.toNumber()
}

if (util.Long !== Long) {
  util.Long = Long as any
  configure()
}
