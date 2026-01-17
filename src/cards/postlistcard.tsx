import React from "react";
import { Box, capitalize, Card, CardContent, Divider, Typography } from "@mui/material";
import DeleteIcon from "@mui/icons-material/Delete";
import EditIcon from "@mui/icons-material/Edit";
import PersonIcon from "@mui/icons-material/Person";

import type { Post } from "../types/entity";

import { showLastUpdated, truncate } from "../utils/formatters";

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
                <Box sx={{ display: "flex", flexDirection: "row", gap: 2 }}>
                    <Box
                        sx={{
                            display: "flex",
                            justifyContent: "center",
                            alignItems: "center",
                            backgroundColor: "lightgrey",
                            borderRadius: 10,
                            height: "4ch",
                            width: "4ch"
                        }}>
                        <PersonIcon />
                    </Box>
                    <Box
                        sx={{
                            display: "flex",
                            justifyContent: "center",
                            alignItems: "center",
                        }}>
                        <Typography>
                            {truncate(capitalize(post.username))} • {showLastUpdated(post.updated_at)}
                        </Typography>
                    </Box>

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
                            <EditIcon
                                onClick={(event: React.MouseEvent<SVGSVGElement, MouseEvent>) => {
                                    event.stopPropagation();
                                    openPostModal(post);
                                }}
                                sx={{
                                    p: "3px",
                                    borderRadius: 10,
                                    color: "blue",
                                    "&:hover": { backgroundColor: "lightgrey" }
                                }}
                            />
                            <DeleteIcon
                                onClick={(event: React.MouseEvent<SVGSVGElement, MouseEvent>) => {
                                    event.stopPropagation();
                                    onDelete(post);
                                }}
                                sx={{
                                    p: "3px",
                                    borderRadius: 10,
                                    color: "red",
                                    "&:hover": { backgroundColor: "lightgrey" }
                                }}
                            />
                        </Box>}
                </Box>

                <Typography
                    sx={{
                        fontSize: "24px",
                        fontWeight: "bold",
                        paddingTop: "8px",
                        display: "-webkit-box",
                        WebkitLineClamp: 2,
                        WebkitBoxOrient: "vertical",
                        textOverflow: "ellipsis",
                        overflow: "hidden",
                        whiteSpace: "pre-wrap"
                    }}>
                    {capitalize(post.title)}
                </Typography>

                <Divider />

                <Typography
                    sx={{
                        paddingTop: "8px",
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
