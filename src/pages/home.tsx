import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Alert, Button, capitalize, Snackbar, Typography } from "@mui/material";
import AddIcon from '@mui/icons-material/Add';
import ArrowBackIosIcon from '@mui/icons-material/ArrowBackIos';

import type { Topic } from "../types/entity";

import EmptyList from "../cards/emptylist";
import TopicListCard from "../cards/topiclistcard";
import TopicModal from "../modals/topic";

import axiosInstance from "../utils/axiosInstance";
import { truncate } from "../utils/formatters";

/**
 * Renders a home page that sends a `GET` request to fetch and displays all topics, a `POST`
 * request to create a new topic, a `PUT` request to update an existing topic and a `DELETE`
 * request to delete a selected topic.
 */
const HomePage: React.FC = () => {
    const [userId] = useState<string>(localStorage.getItem("token") ?? "");
    const [topic, setTopic] = useState<Topic | null>(null);
    const [topiclist, setTopiclist] = useState<Topic[]>([]);
    const [isError, setIsError] = useState<boolean>(false);
    const [message, setMessage] = useState<string>("");
    const [openModal, setOpenModal] = useState<boolean>(false);
    const [isUpdate, setIsUpdate] = useState<boolean>(false);
    const [openSnackBar, setOpenSnackBar] = useState<boolean>(false);
    const navigate = useNavigate();

    const getAllTopics = () => {
        axiosInstance.get("/topics")
            .then(res => {
                if (res.data) {
                    setIsError(false);
                    setTopiclist(res.data.payload?.data);
                }
            })
            .catch(err => {
                console.error("unable to get all topics: " + err);
                setMessage(err);
                setIsError(true);
                setOpenSnackBar(true);
            });
    };

    const handleBack = () => {
        localStorage.clear();
        navigate("/");
    };

    const handleClick = (id: number) => {
        navigate(`/home/${id}`);
    };

    const closeModal = () => {
        setOpenModal(false);
        setIsUpdate(false);
    };

    const openTopicModal = (t: Topic) => {
        setIsUpdate(true);
        setTopic(t);
        setOpenModal(true);
    };

    const closeSnackBar = () => {
        setOpenSnackBar(false);
        setMessage("");
    };

    /**
     * Creates the new topic. The title is converted and stored in all lowercase in the datebase.
     */
    const onCreate = (title: string) => {
        setOpenSnackBar(false);
        setMessage("");
        setIsError(false);

        if (userId == "") {
            setIsError(true);
            setMessage("system error: userId misssing");
            setOpenSnackBar(true);
            return;
        }

        axiosInstance.post("/topics", {
            userId: Number.parseInt(userId, 10), title: title.toLocaleLowerCase()
        })
            .then(res => {
                if (res.data) {
                    setMessage("Created " + capitalize(title));
                    setIsError(false);
                    getAllTopics();
                }
            })
            .catch(err => {
                console.error("unable to add new topic: " + err);
                setMessage(err);
                setIsError(true);
            })
            .finally(() => {
                setOpenSnackBar(true);
                closeModal();
            });
    };

    /**
     * Updates the selected topic. The title is converted and stored in all lowercase in the datebase.
     * Only the author is able to update the topic.
     */
    const onUpdate = (topicId: number, topicUserId: number, title: string) => {
        setOpenSnackBar(false);
        setMessage("");
        setIsError(false);

        if (userId == "") {
            setIsError(true);
            setMessage("system error: userId misssing");
            setOpenSnackBar(true);
            return;
        }

        if (userId != topicUserId.toString()) {
            setIsError(true);
            setMessage("Not creator. Unable to update.");
            setOpenSnackBar(true);
            return;
        }

        axiosInstance.put(`/topics/${topicId}`, { title: title.toLocaleLowerCase() })
            .then(res => {
                if (res.data) {
                    setMessage("Updated " + capitalize(title));
                    setIsError(false);
                    getAllTopics();
                }
            })
            .catch(err => {
                console.error("unable to update topic: " + err);
                setMessage(err);
                setIsError(true);
            })
            .finally(() => {
                setOpenSnackBar(true);
                closeModal();
            });
    };

    /**
     * Deletes the selected topic. Only the author is able to delete the topic.
     */
    const onDelete = (t: Topic) => {
        setOpenSnackBar(false);
        setMessage("");
        setIsError(false);

        if (userId == "") {
            setIsError(true);
            setMessage("system error: userId misssing");
            setOpenSnackBar(true);
            return;
        }

        if (userId != t.user_id.toString()) {
            setIsError(true);
            setMessage("Not creator. Unable to delete.");
            setOpenSnackBar(true);
            return;
        }

        axiosInstance.delete(`/topics/${t.topic_id}`)
            .then(res => {
                if (res.data) {
                    setIsError(false);
                    setMessage("Deleted " + capitalize(t.title));
                    getAllTopics();
                }
            })
            .catch(err => {
                console.error("unable to delete topic: " + err);
                setMessage(err);
                setIsError(true);
            })
            .finally(() => setOpenSnackBar(true));
    };

    useEffect(() => {
        setIsError(false);
        setMessage("");
        getAllTopics();
    }, []);

    return (
        <div
            style={{
                position: "relative",
                display: "flex",
                alignItems: "center",
                flexDirection: "column",
            }}>
            <Button
                variant="outlined"
                size="large"
                onClick={handleBack}
                sx={{
                    position: "absolute",
                    top: 48,
                    left: 80,
                    borderRadius: 3,
                    fontSize: "20px",
                    "&:hover": { backgroundColor: "#5aacfdff", color: "white" }
                }}>
                <ArrowBackIosIcon sx={{ fontSize: "20px" }} />
                Back
            </Button>

            <Typography variant="h4" gutterBottom sx={{ fontWeight: "bold", marginTop: "32px" }}>
                Click the topics below to share your insights.
            </Typography>
            <Typography variant="h6" sx={{ fontWeight: "bold", marginBottom: "20px" }}>
                Or click the 'Add' button to add a new topic.
            </Typography>

            <Button
                variant="outlined"
                size="large"
                onClick={() => setOpenModal(true)}
                sx={{
                    position: "absolute",
                    top: 48,
                    right: 80,
                    borderRadius: 3,
                    fontSize: "20px",
                    "&:hover": { backgroundColor: "#5aacfdff", color: "white" }
                }}>
                <AddIcon sx={{ display: "flex", fontSize: "24px", mr: 0.25 }} />
                Add
            </Button>

            {(topiclist == null || topiclist.length == 0)
                ? <EmptyList entity="topic" />
                : topiclist.map((t: Topic) => (
                    <TopicListCard
                        key={t.topic_id}
                        topic={t}
                        isUser={userId == t.user_id.toString()}
                        handleClick={handleClick}
                        openTopicModal={openTopicModal}
                        onDelete={onDelete}
                    />
                ))
            }

            <div style={{ margin: "20px" }} />

            <TopicModal
                open={openModal}
                close={closeModal}
                topic={topic}
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

export default HomePage;
