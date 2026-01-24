import React, { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { Alert, capitalize, Snackbar, Typography } from "@mui/material";

import type { Comment, Post } from "../types/entity";

import EmptyList from "../cards/emptylist";
import PostCard from "../cards/postcard";
import BackButton from "../components/backbutton";
import PostModal from "../modals/post";

import axiosInstance from "../utils/axiosInstance";
import { truncate } from "../utils/formatters";

/**
 * Renders a post page that sends a `GET` request to fetch and display the specific post and all
 * its comments, a `POST` request to create a new comment, a `PUT` request to update an existing
 * post or comment and a `DELETE` request to delete a selected comment. It also sends a `POST`
 * request to like and dislike a comment, and a `DELETE` request to remove the user's vote.
 */
const PostPage: React.FC = () => {
    const topicId: string = useParams().topicId ?? "";
    const postId: string = useParams().postId ?? "";
    const [post, setPost] = useState<Post | null>(null);
    const [commentList, setCommentList] = useState<Comment[]>([]);
    const [isError, setIsError] = useState<boolean>(false);
    const [message, setMessage] = useState<string>("");
    const [openModal, setOpenModal] = useState<boolean>(false);
    const [openSnackBar, setOpenSnackBar] = useState<boolean>(false);
    const navigate = useNavigate();

    const getPost = (topicId: string, postId: string) => {
        axiosInstance.get(`/api/posts/${topicId}/${postId}`)
            .then(res => {
                if (res.data) {
                    setIsError(false);
                    setPost(res.data.payload?.data);
                }
            })
            .catch(err => {
                console.error("unable to get post: " + err);
                setMessage(err);
                setIsError(true);
                setOpenSnackBar(true);
            })
    };

    const getAllComments = (topicId: string, postId: string) => {
        axiosInstance.get(`/api/comments/all/${topicId}/${postId}`)
            .then(res => {
                if (res.data) {
                    setIsError(false);
                    setCommentList(res.data.payload?.data);
                }
            })
            .catch(err => {
                console.error("unable to get all comments: " + err);
                setMessage(err);
                setIsError(true);
                setOpenSnackBar(true);
            });
    };

    const handleBack = () => {
        if (!topicId) {
            setIsError(true);
            setMessage("system error: topicId missing");
            setOpenSnackBar(true);
            return;
        }
        navigate(`/home/${topicId}`);
    };

    const closeSnackBar = () => {
        setOpenSnackBar(false);
        setMessage("");
    };

    const openPostModal = (p: Post) => {
        setPost(p);
        setOpenModal(true);
    };

    /**
     * Updates the selected post. Only the author is able to update the post.
     */
    const updatePost = (postId: number, title: string, description: string) => {
        setOpenSnackBar(false);
        setMessage("");
        setIsError(false);

        axiosInstance.put(`/api/posts/${postId}`, {
            title: title.trim(),
            description: description.trim()
        })
            .then(res => {
                if (res.data) {
                    setMessage("Updated " + capitalize(title.trim()));
                    setIsError(false);
                    getPost(topicId, postId.toString());
                }
            })
            .catch(err => {
                console.error("unable to update post: " + err);
                setMessage(err);
                setIsError(true);
            })
            .finally(() => {
                setOpenSnackBar(true);
                setOpenModal(false);
            });
    };

    const onCreate = (description: string) => {
        setOpenSnackBar(false);
        setMessage("");
        setIsError(false);

        axiosInstance.post("/api/comments", {
            postId: Number(postId),
            description: description.trim()
        })
            .then(res => {
                if (res.data) {
                    setMessage("Comment successfully!");
                    setIsError(false);
                    getPost(topicId, postId);
                    getAllComments(topicId, postId);
                }
            })
            .catch(err => {
                console.error("unable to add new comment: " + err);
                setMessage(err);
                setIsError(true);
            })
            .finally(() => setOpenSnackBar(true));
    };

    /**
     * Updates the selected comment. Only the author is able to update the comment.
     */
    const onUpdate = (commentId: number, commentPostId: number, description: string) => {
        setOpenSnackBar(false);
        setMessage("");
        setIsError(false);

        axiosInstance.put(`/api/comments/${commentId}`, {
            postId: commentPostId,
            description: description.trim()
        })
            .then(res => {
                if (res.data) {
                    setMessage("Comment updated!");
                    setIsError(false);
                    getPost(topicId, postId);
                    getAllComments(topicId, postId);
                }
            })
            .catch(err => {
                console.error("unable to update comment: " + err);
                setMessage(err);
                setIsError(true);
            })
            .finally(() => setOpenSnackBar(true));
    };

    /**
     * Deletes the selected comment. Only the author is able to delete the comment.
     */
    const onDelete = (c: Comment) => {
        setOpenSnackBar(false);
        setMessage("");
        setIsError(false);

        axiosInstance.delete(`/api/comments/${c.comment_id}`)
            .then(res => {
                if (res.data) {
                    setIsError(false);
                    setMessage("Comment deleted!");
                    getPost(topicId, postId);
                    getAllComments(topicId, postId);
                }
            })
            .catch(err => {
                console.error("unable to delete comment: " + err);
                setMessage(err);
                setIsError(true);
            })
            .finally(() => setOpenSnackBar(true));
    };

    /**
     * Sets a like on the entity (post or comment).
     */
    const onLike = (id: number, entityType: string) => {
        setOpenSnackBar(false);
        setMessage("");
        setIsError(false);

        axiosInstance.post(`/api/${entityType}s/${id}/likes`)
            .then(res => {
                if (res.data) {
                    entityType == "comment"
                        ? getAllComments(topicId, postId)
                        : getPost(topicId, postId);
                    setIsError(false);
                    setMessage("Liked!");
                }
            })
            .catch(err => {
                console.error("unable to like " + entityType + ": " + err);
                setIsError(true);
                setMessage(err);
            })
            .finally(() => setOpenSnackBar(true));
    };

    /**
     * Sets a dislike on the entity (post or comment).
     */
    const onDislike = (id: number, entityType: string) => {
        setOpenSnackBar(false);
        setMessage("");
        setIsError(false);

        axiosInstance.post(`/api/${entityType}s/${id}/dislikes`)
            .then(res => {
                if (res.data) {
                    entityType == "comment"
                        ? getAllComments(topicId, postId)
                        : getPost(topicId, postId);
                    setIsError(false);
                    setMessage("Disliked!");
                }
            })
            .catch(err => {
                console.error("unable to dislike " + entityType + ": " + err);
                setIsError(true);
                setMessage(err);
            })
            .finally(() => setOpenSnackBar(true));
    };

    /**
     * Removes the vote on the entity (post or comment).
     */
    const onRemoveVote = (id: number, entityType: string) => {
        setOpenSnackBar(false);
        setMessage("");
        setIsError(false);

        axiosInstance.delete(`/api/${entityType}s/${id}/remove`)
            .then(res => {
                if (res.data) {
                    entityType == "comment"
                        ? getAllComments(topicId, postId)
                        : getPost(topicId, postId);
                    setIsError(false);
                    setMessage("Removed");
                }
            })
            .catch(err => {
                console.error("unable to remove vote on " + entityType + ": " + err);
                setIsError(true);
                setMessage(err);
            })
            .finally(() => setOpenSnackBar(true));
    };

    useEffect(() => {
        setIsError(false);
        setMessage("");
        if (postId) {
            getPost(topicId, postId);
            getAllComments(topicId, postId);
        }
    }, [postId]);

    return (
        <div
            style={{
                position: "relative",
                display: "flex",
                alignItems: "center",
                flexDirection: "column",
            }}>
            <BackButton handleBack={handleBack} />

            <Typography variant="h4" gutterBottom sx={{ fontWeight: "bold", mt: "32px" }}>
                Post: {truncate(capitalize(post?.title ?? ""), 42)}
            </Typography>

            {post
                ? <PostCard
                    post={post}
                    commentList={commentList}
                    openPostModal={openPostModal}
                    onCreate={onCreate}
                    onUpdate={onUpdate}
                    onDelete={onDelete}
                    onLike={onLike}
                    onDislike={onDislike}
                    onRemoveVote={onRemoveVote}
                />
                : <EmptyList entity="post" />
            }

            <div style={{ margin: "20px" }} />

            <PostModal
                open={openModal}
                close={() => setOpenModal(false)}
                post={post}
                isUpdate={true}
                onCreate={() => { }}
                onUpdate={updatePost}
            />

            <Snackbar open={openSnackBar} autoHideDuration={5000} onClose={closeSnackBar}>
                <Alert
                    onClose={closeSnackBar}
                    severity={isError ? "error" : "success"}
                    variant="filled">
                    {truncate(message, 30)}
                </Alert>
            </Snackbar>
        </div>
    );
};

export default PostPage;
