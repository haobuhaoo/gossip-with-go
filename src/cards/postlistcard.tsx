import React from "react";
import { Box, capitalize, Card, CardContent, Divider, Typography } from "@mui/material";

import type { Post } from "../types/entity";

import AvatarIcon from "../components/avataricon";
import DeleteButton from "../components/deletebutton";
import DisplayAuthor from "../components/displayauthor";
import EditButton from "../components/editbutton";

type Props = {
    /**
     * Post to be displayed.
     */
    post: Post;

    /**
     * True if the user is the author of the post.
     */
    isUser: boolean;

    /**
     * Function that passes selected post id to parent component.
     */
    handleClick: (id: number) => void;

    /**
     * Function that passes post to be updated to parent component.
     */
    openPostModal: (t: Post) => void;

    /**
     * Function that passes post to be deleted to parent component.
     */
    onDelete: (t: Post) => void;
}

/**
 * Renders a single post inside a card, which calls `handleClick`, `openPostModal` and
 * `onDelete` when clicked.
 */
const PostListCard: React.FC<Props> = ({ post, isUser, handleClick, openPostModal, onDelete }) => {
    return (
        <Card
            onClick={() => handleClick(post.post_id)}
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
                <Box sx={{ display: "flex", flexDirection: "row", gap: 2, mx: "8px" }}>
                    <AvatarIcon username={post.username} />
                    <DisplayAuthor entity={post} />

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
                            <EditButton<Post> entity={post} updateEntity={openPostModal} />
                            <DeleteButton<Post> entity={post} onDelete={onDelete} />
                        </Box>}
                </Box>

                <Typography
                    sx={{
                        fontSize: "24px",
                        fontWeight: "bold",
                        paddingTop: "8px",
                        mx: "8px",
                        display: "-webkit-box",
                        WebkitLineClamp: 2,
                        WebkitBoxOrient: "vertical",
                        textOverflow: "ellipsis",
                        overflow: "hidden",
                        whiteSpace: "pre-wrap"
                    }}>
                    {capitalize(post.title)}
                </Typography>

                <Divider sx={{ mx: "8px" }} />

                <Typography
                    sx={{
                        paddingTop: "8px",
                        mx: "8px",
                        display: "-webkit-box",
                        WebkitLineClamp: 6,
                        WebkitBoxOrient: "vertical",
                        textOverflow: "ellipsis",
                        overflow: "hidden",
                        whiteSpace: "pre-wrap"
                    }}>
                    {capitalize(post.description)}
                </Typography>
            </CardContent>
        </Card>
    )
}

export default PostListCard;
