import React, { useState } from "react";
import {
    Box, Button, capitalize, Card, CardContent, Divider, InputAdornment, TextField, Typography
} from "@mui/material";

import type { Comment, Post } from "../types/entity";

import AvatarIcon from "../components/avataricon";
import DisplayAuthor from "../components/displayauthor";
import EditButton from "../components/editbutton";

import EmptyList from "./emptylist";
import CommentListCard from "./commentlistcard";

type Props = {
    /**
     * Post to be displayed.
     */
    post: Post;

    /**
     * Comments to be displayed.
     */
    commentList: Comment[];

    /**
     * Function that passes post to be updated to parent component.
     */
    openPostModal: (t: Post) => void;

    /**
     * Function that passes `description` to be created to parent component.
     */
    onCreate: (description: string) => void;

    /**
     * Function that passes the updated `description` to parent component, along with its
     * `commentId`, `commentPostId` and `commentUserId`.
     */
    onUpdate: (
        commentId: number, commentPostId: number, commentUserId: number, description: string) => void;

    /**
     * Function that passes comment to be deleted to parent component.
     */
    onDelete: (c: Comment) => void;
}

/**
 * Renders a full post with its comments, which allows the creation, edit and deletion of comments.
 */
const PostCard: React.FC<Props> = ({
    post, commentList, openPostModal, onCreate, onUpdate, onDelete }) => {
    const userId: string = localStorage.getItem("token") ?? "";
    const [comment, setComment] = useState<string>("");

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setComment(event.target.value);
    };

    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        onCreate(comment);
        setComment("");
    };

    return (
        <Card
            sx={{
                width: "60vw",
                margin: "8px",
                border: "2px solid",
                borderColor: "divider",
                borderRadius: 5,
                boxShadow: "0 1px 3px rgba(0, 0, 0, 0.12), 0 2px 4px rgba(0, 0, 0, 0.08)",
                cursor: "pointer",
                "&:hover": { boxShadow: "0 6px 12px rgba(0, 0, 0, 0.12)" },
            }}>
            <CardContent sx={{ display: "flex", flexDirection: "column" }}>
                <Box sx={{ display: "flex", flexDirection: "row", gap: 2, mx: "16px" }}>
                    <AvatarIcon username={post.username} />
                    <DisplayAuthor entity={post} />

                    {userId == post.user_id.toString() &&
                        <Box
                            sx={{
                                display: "flex",
                                alignItems: "center",
                                marginLeft: "auto",
                                fontSize: "24px",
                                paddingRight: "8px",
                                gap: "4px",
                            }}>
                            <EditButton<Post> entity={post} updateEntity={openPostModal} />
                        </Box>}
                </Box>

                <Typography
                    sx={{
                        fontSize: "24px",
                        fontWeight: "bold",
                        paddingTop: "8px",
                        mx: "16px",
                        whiteSpace: "pre-wrap"
                    }}>
                    {capitalize(post.title)}
                </Typography>

                <Divider sx={{ mx: "16px" }} />

                <Typography sx={{ paddingTop: "8px", mx: "16px", whiteSpace: "pre-wrap" }}>
                    {capitalize(post.description)}
                </Typography>

                <Box
                    component="form"
                    onSubmit={handleSubmit}
                    sx={{ mt: 3, mx: 1 }}>
                    <TextField
                        id="comment"
                        value={comment}
                        placeholder="Join the conversation"
                        autoComplete="off"
                        multiline
                        rows={2}
                        onChange={handleChange}
                        fullWidth
                        sx={{ "& .MuiOutlinedInput-root": { borderRadius: 10, pl: 4 } }}
                        slotProps={{
                            input: {
                                endAdornment: (
                                    <InputAdornment position="end">
                                        <Button type="submit">Add</Button>
                                    </InputAdornment>
                                ),
                            }
                        }}
                    />
                </Box>

                {(commentList == null || commentList.length == 0)
                    ? <EmptyList entity="comment" />
                    : commentList.map((c: Comment) => (
                        <CommentListCard
                            key={c.comment_id}
                            comment={c}
                            isUser={userId == c.user_id.toString()}
                            onUpdate={onUpdate}
                            onDelete={onDelete}
                        />
                    ))
                }
            </CardContent>
        </Card>
    )
}

export default PostCard;
