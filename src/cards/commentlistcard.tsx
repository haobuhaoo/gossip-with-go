import React, { useState } from "react";
import { Box, Button, capitalize, Card, CardContent, Divider, TextField, Typography } from "@mui/material";

import type { Comment } from "../types/entity";

import AvatarIcon from "../components/avataricon";
import DeleteButton from "../components/deletebutton";
import DisplayAuthor from "../components/displayauthor";
import EditButton from "../components/editbutton";
import VoteButton from "../components/votebutton";

type Props = {
    /**
     * Comment to be displayed.
     */
    comment: Comment;

    /**
     * True if the user is the author of the comment.
     */
    isUser: boolean;

    /**
     * Function that passes the updated `description` to parent component, along with its
     * `commentId` and `commentPostId``.
     */
    onUpdate: (commentId: number, commentPostId: number, description: string) => void;

    /**
     * Function that passes comment to be deleted to parent component.
     */
    onDelete: (c: Comment) => void;

    /**
     * Function that passes the commentId number and entityType back to the parent component to indicate
     * a like for the comment.
     */
    onLike: (id: number, entity: string) => void;

    /**
     * Function that passes the commentId number and entityType back to the parent component to indicate
     * a dislike for the comment.
     */
    onDislike: (id: number, entity: string) => void;

    /**
     * Function that passes the commentId number and entityType back to the parent component to remove
     * user's vote for the comment.
     */
    onRemoveVote: (id: number, entity: string) => void;
}

/**
 * Renders a single comment inside a card, which calls `onUpdate`, `onDelete`, `onLike`, `onDislike`
 * and `onRemoveVote` when clicked.
 */
const CommentListCard: React.FC<Props> = ({
    comment, isUser, onUpdate, onDelete, onLike, onDislike, onRemoveVote }) => {
    const [desc, setDesc] = useState<string>(comment.description ?? "");
    const [isUpdate, setIsUpdate] = useState<boolean>(false);

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setDesc(event.target.value);
    };

    const handleBlur = (event: React.FocusEvent<HTMLInputElement>) => {
        if ((event.relatedTarget as HTMLElement)?.id === "update") return;
        setIsUpdate(false);
    };

    const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.stopPropagation();
        onUpdate(comment.comment_id, comment.post_id, desc);
        setIsUpdate(false);
    };

    return (
        <Card sx={{ boxShadow: 0 }}>
            <CardContent sx={{ display: "flex", flexDirection: "column" }}>
                <Box sx={{ display: "flex", gap: 2 }}>
                    <AvatarIcon username={comment.username} />
                    <DisplayAuthor entity={comment} />

                    {isUser &&
                        <Box
                            sx={{
                                display: "flex",
                                alignItems: "center",
                                marginLeft: "auto",
                                fontSize: "24px",
                                paddingRight: "8px",
                                gap: "4px",
                            }}>
                            {isUpdate ?
                                <Button
                                    id="update"
                                    onClick={handleClick}
                                    sx={{
                                        fontWeight: "bold",
                                        fontSize: "16px",
                                        borderRadius: 2,
                                        "&:hover": { backgroundColor: "#5aacfdff", color: "white" }
                                    }}>
                                    Update
                                </Button>
                                : <Box>
                                    <EditButton<Comment>
                                        entity={comment}
                                        updateEntity={(_) => setIsUpdate(true)}
                                    />
                                    <DeleteButton<Comment>
                                        entity={comment}
                                        onDelete={onDelete}
                                    />
                                </Box>}
                        </Box>}
                </Box>

                <Box sx={{ display: "flex", marginTop: "8px" }}>
                    <Divider orientation="vertical" flexItem sx={{ mx: 2 }} />

                    <Box>
                        {isUpdate ?
                            <TextField
                                id="comment"
                                value={desc}
                                autoComplete="off"
                                autoFocus
                                multiline
                                fullWidth
                                onChange={handleChange}
                                onBlur={handleBlur}
                            />
                            : <Typography sx={{ whiteSpace: "pre-wrap" }}>
                                {capitalize(comment.description)}
                            </Typography>}


                        <Box sx={{ display: "flex", mt: 1, gap: 1 }}>
                            {["likes", "dislikes"].map((s: string) => (
                                <VoteButton
                                    type={s == "likes" ? "like" : "dislike"}
                                    vote={comment.user_vote}
                                    voteCount={s == "likes" ? comment.likes : comment.dislikes}
                                    id={comment.comment_id}
                                    entityType="comment"
                                    onLike={onLike}
                                    onDislike={onDislike}
                                    onRemoveVote={onRemoveVote}
                                />
                            ))}
                        </Box>
                    </Box>
                </Box>
            </CardContent>
        </Card>
    )
}

export default CommentListCard;
