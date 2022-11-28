import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateComment } from "./types/blog/blog/tx";
import { MsgCreatePost } from "./types/blog/blog/tx";
import { MsgDeleteComment } from "./types/blog/blog/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/blog.blog.MsgCreateComment", MsgCreateComment],
    ["/blog.blog.MsgCreatePost", MsgCreatePost],
    ["/blog.blog.MsgDeleteComment", MsgDeleteComment],
    
];

export { msgTypes }