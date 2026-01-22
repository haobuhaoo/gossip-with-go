import React, { useEffect, useState } from "react";
import { Box, Button, Dialog, DialogTitle, TextField } from "@mui/material";

import type { Topic } from "../types/entity";

import CloseModalButton from "../components/closemodalbutton";

type Props = {
    /**
     * True if modal is open.
     */
    open: boolean;

    /**
     * Function that closes the modal.
     */
    close: () => void;

    /**
     * Topic to be updated, or null if creating a new topic.
     */
    topic: Topic | null;

    /**
     * True if updating an existing topic.
     */
    isUpdate: boolean;

    /**
     * Function that passes the `newTopic` to be created to parent component.
     */
    onCreate: (newTopic: string) => void;

    /**
     * Function that passes the updated `newTopic` to parent component, along with its `topicId`.
     */
    onUpdate: (topicId: number, newTopic: string) => void;
}

/**
 * Renders a form modal that calls `onCreate` and `onUpdate` when button is clicked.
 */
const TopicModal: React.FC<Props> = ({ open, close, topic, isUpdate, onCreate, onUpdate }) => {
    const [newTopic, setNewTopic] = useState<string>("");
    const [error, setError] = useState<string>(" ");

    const handleClose = () => {
        setNewTopic("");
        close();
    };

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setNewTopic(event.target.value);
    };

    /**
     * Sends the title to create if `isUpdate` is false or update the topic if `isUpdate`
     * is true.
     */
    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setError(" ");

        if (isUpdate) {
            if (topic == null) {
                setError("system error: topic missing");
                return;
            }
            onUpdate(topic.topic_id, newTopic);
            setNewTopic("");
            setError(" ");
            return;
        }
        onCreate(newTopic);
        setNewTopic("");
        setError(" ");
    };

    useEffect(() => {
        if (error) setTimeout(() => setError(" "), 5000);
    }, [error]);

    useEffect(() => {
        if (isUpdate) setNewTopic(topic?.title || "");
    }, [open, topic]);

    return (
        <Dialog open={open} onClose={() => handleClose()} disableRestoreFocus>
            <DialogTitle
                sx={{
                    display: "flex",
                    justifyContent: "center",
                    alignItems: "center",
                    fontWeight: "bold",
                    fontSize: "28px"
                }}>
                {isUpdate ? "Enter the new topic" : "Enter a new topic."}
            </DialogTitle>
            <CloseModalButton close={close} />

            <Box
                component="form"
                onSubmit={handleSubmit}
                sx={{
                    display: "flex",
                    flexDirection: "column",
                    justifyContent: "center",
                    alignItems: "center",
                    marginX: "20px",
                    marginBottom: "20px",
                    "& .MuiTextField-root": { mt: 2, mb: 1, mx: 3, width: "40ch" },
                }}>
                <TextField
                    required
                    id="newTopic"
                    label="New Topic"
                    value={newTopic}
                    placeholder="topic"
                    autoComplete="off"
                    autoFocus
                    onChange={handleChange}
                    error={error != " "}
                    helperText={error}
                />
                <Button
                    variant="outlined"
                    type="submit"
                    sx={{
                        fontWeight: "bold",
                        fontSize: "16px",
                        borderRadius: 2,
                        "&:hover": { backgroundColor: "#5aacfdff", color: "white" }
                    }}>
                    {isUpdate ? "Update" : "Add"}
                </Button>
            </Box>
        </Dialog>
    )
};

export default TopicModal;
