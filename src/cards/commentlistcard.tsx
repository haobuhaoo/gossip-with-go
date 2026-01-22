import React, { useState } from "react";
import { Box, Button, capitalize, Card, CardContent, Divider, TextField, Typography } from "@mui/material";

import type { Comment } from "../types/entity";

import AvatarIcon from "../components/avataricon";
import DeleteButton from "../components/deletebutton";
import DisplayAuthor from "../components/displayauthor";
import EditButton from "../components/editbutton";

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
}

/**
 * Renders a single comment inside a card, which calls `onUpdate` and `onDelete`when clicked.
 */
const CommentListCard: React.FC<Props> = ({ comment, isUser, onUpdate, onDelete }) => {
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
                <Box sx={{ display: "flex", flexDirection: "row", gap: 2 }}>
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
                                    <EditButton<Comment> entity={comment} updateEntity={(_) => setIsUpdate(true)} />
                                    <DeleteButton<Comment> entity={comment} onDelete={onDelete} />
                                </Box>}
                        </Box>}
                </Box>

                <Box sx={{ display: "flex", flexDirection: "row", marginTop: "8px" }}>
                    <Divider orientation="vertical" flexItem sx={{ mx: 2 }} />

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
                </Box>
            </CardContent>
        </Card>
    )
}

export default CommentListCard;
