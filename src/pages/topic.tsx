import React, { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { Alert, capitalize, Snackbar, Typography } from "@mui/material";

import type { Post, Topic } from "../types/entity";

import { useAuth } from "../context/authcontext";

import EmptyList from "../cards/emptylist";
import PostListCard from "../cards/postlistcard";
import AddButton from "../components/addbutton";
import BackButton from "../components/backbutton";
import SearchBar from "../components/searchbar";
import PostModal from "../modals/post";

import axiosInstance from "../utils/axiosInstance";
import { truncate } from "../utils/formatters";

/**
 * Renders a topic page that sends a `GET` request to fetch and displays all posts under the
 * specific topic, a `POST` request to create a new post, a `PUT` request to update an existing
 * post and a `DELETE` request to delete a selected post. It also sends a `POST` request to like
 * and dislike a post, and a `DELETE` request to remove the user's vote.
 */
const TopicPage: React.FC = () => {
    const topicId: string = useParams().topicId ?? "";
    const [topic, setTopic] = useState<Topic>();
    const [post, setPost] = useState<Post | null>(null);
    const [postlist, setPostlist] = useState<Post[]>([]);
    const [isError, setIsError] = useState<boolean>(false);
    const [message, setMessage] = useState<string>("");
    const [openModal, setOpenModal] = useState<boolean>(false);
    const [isUpdate, setIsUpdate] = useState<boolean>(false);
    const [openSnackBar, setOpenSnackBar] = useState<boolean>(false);
    const { auth } = useAuth();
    const navigate = useNavigate();

    const getTopic = (id: string) => {
        axiosInstance.get(`/api/topics/${id}`)
            .then(res => {
                if (res.data) {
                    setIsError(false);
                    setTopic(res.data.payload?.data);
                }
            })
            .catch(err => {
                console.error("unable to get topic: " + err);
                setMessage(err);
                setIsError(true);
                setOpenSnackBar(true);
            })
    };

    const getAllPosts = (id: string) => {
        axiosInstance.get(`/api/posts/all/${id}`)
            .then(res => {
                if (res.data) {
                    setIsError(false);
                    setPostlist(res.data.payload?.data);
                }
            })
            .catch(err => {
                console.error("unable to get all posts: " + err);
                setMessage(err);
                setIsError(true);
                setOpenSnackBar(true);
            });
    };

    const handleBack = () => {
        navigate("/home");
    };

    const handleClick = (id: number) => {
        navigate(`/home/${topicId}/${id}`);
    };

    const closeModal = () => {
        setOpenModal(false);
        setIsUpdate(false);
    };

    const openPostModal = (p: Post) => {
        setIsUpdate(true);
        setPost(p);
        setOpenModal(true);
    };

    const closeSnackBar = () => {
        setOpenSnackBar(false);
        setMessage("");
    };

    const onCreate = (title: string, description: string) => {
        setOpenSnackBar(false);
        setMessage("");
        setIsError(false);

        axiosInstance.post("/api/posts", {
            topicId: Number(topicId),
            title: title.trim(),
            description: description.trim()
        })
            .then(res => {
                if (res.data) {
                    setMessage("Created " + capitalize(title.trim()));
                    setIsError(false);
                    getAllPosts(topicId);
                }
            })
            .catch(err => {
                console.error("unable to add new post: " + err);
                setMessage(err);
                setIsError(true);
            })
            .finally(() => {
                setOpenSnackBar(true);
                closeModal();
            });
    };

    /**
     * Updates the selected post. Only the author is able to update the post.
     */
    const onUpdate = (
        postId: number, title: string, description: string) => {
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
                    getAllPosts(topicId);
                }
            })
            .catch(err => {
                console.error("unable to update post: " + err);
                setMessage(err);
                setIsError(true);
            })
            .finally(() => {
                setOpenSnackBar(true);
                closeModal();
            });
    };

    /**
     * Deletes the selected post. Only the author is able to delete the post.
     */
    const onDelete = (p: Post) => {
        setOpenSnackBar(false);
        setMessage("");
        setIsError(false);

        axiosInstance.delete(`/api/posts/${p.post_id}`)
            .then(res => {
                if (res.data) {
                    setIsError(false);
                    setMessage("Deleted " + capitalize(p.title.trim()));
                    getAllPosts(topicId);
                }
            })
            .catch(err => {
                console.error("unable to delete post: " + err);
                setMessage(err);
                setIsError(true);
            })
            .finally(() => setOpenSnackBar(true));
    };

    const onLike = (id: number) => {
        setOpenSnackBar(false);
        setMessage("");
        setIsError(false);

        axiosInstance.post(`/api/posts/${id}/likes`)
            .then(res => {
                if (res.data) {
                    getAllPosts(topicId);
                    setIsError(false);
                    setMessage("Liked!");
                }
            })
            .catch(err => {
                console.error("unable to like post: " + err);
                setIsError(true);
                setMessage(err);
            })
            .finally(() => setOpenSnackBar(true));
    };

    const onDislike = (id: number) => {
        setOpenSnackBar(false);
        setMessage("");
        setIsError(false);

        axiosInstance.post(`/api/posts/${id}/dislikes`)
            .then(res => {
                if (res.data) {
                    getAllPosts(topicId);
                    setIsError(false);
                    setMessage("Disliked!");
                }
            })
            .catch(err => {
                console.error("unable to dislike post: " + err);
                setIsError(true);
                setMessage(err);
            })
            .finally(() => setOpenSnackBar(true));
    };

    const onRemoveVote = (id: number) => {
        setOpenSnackBar(false);
        setMessage("");
        setIsError(false);

        axiosInstance.delete(`/api/posts/${id}/remove`)
            .then(res => {
                if (res.data) {
                    getAllPosts(topicId);
                    setIsError(false);
                    setMessage("Removed");
                }
            })
            .catch(err => {
                console.error("unable to remove vote on post: " + err);
                setIsError(true);
                setMessage(err);
            })
            .finally(() => setOpenSnackBar(true));
    };

    useEffect(() => {
        setIsError(false);
        setMessage("");
        if (topicId) {
            getTopic(topicId);
            getAllPosts(topicId);
        }
    }, [topicId]);

    return (
        <div
            style={{
                position: "relative",
                display: "flex",
                alignItems: "center",
                flexDirection: "column",
            }}>
            <BackButton handleBack={handleBack} />

            <Typography variant="h4" gutterBottom sx={{ fontWeight: "bold", marginTop: "32px" }}>
                Topic: {truncate(capitalize(topic?.title ?? ""), 42)}
            </Typography>
            <Typography variant="h5" gutterBottom sx={{ fontWeight: "bold" }}>
                Click the posts below to read more or give your comments.
            </Typography>
            <Typography variant="h6" sx={{ fontWeight: "bold", marginBottom: "20px" }}>
                Or click the 'Add' button to add a new post.
            </Typography>

            <AddButton setOpenModal={setOpenModal} />

            <SearchBar<Post>
                isTopic={false}
                topicId={topicId}
                setEntity={setPostlist}
                setMessage={setMessage}
                setIsError={setIsError}
                setOpenSnackBar={setOpenSnackBar}
            />

            {(postlist == null || postlist.length == 0)
                ? <EmptyList entity="post" />
                : postlist.map((p: Post) => (
                    <PostListCard
                        key={p.post_id}
                        post={p}
                        isUser={auth.userId == p.user_id.toString()}
                        handleClick={handleClick}
                        openPostModal={openPostModal}
                        onDelete={onDelete}
                        onLike={onLike}
                        onDislike={onDislike}
                        onRemoveVote={onRemoveVote}
                    />
                ))
            }

            <div style={{ margin: "20px" }} />

            <PostModal
                open={openModal}
                close={closeModal}
                post={post}
                isUpdate={isUpdate}
                onCreate={onCreate}
                onUpdate={onUpdate}
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

export default TopicPage;
