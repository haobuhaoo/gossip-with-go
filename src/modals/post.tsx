import React, { useEffect, useState } from "react";
import { Box, Button, Dialog, DialogTitle, TextField } from "@mui/material";

import type { Post } from "../types/entity";

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
     * Post to be updated, or null if creating a new post.
     */
    post: Post | null;

    /**
     * True if updating an existing post.
     */
    isUpdate: boolean;

    /**
     * Function that passes the `title` and `description` to be created to parent component.
     */
    onCreate: (title: string, description: string) => void;

    /**
     * Function that passes the updated `title` and `description` to parent component,
     * along with its `postId`.
     */
    onUpdate: (postId: number, title: string, description: string) => void;
}

/**
 * Renders a form modal that calls `onCreate` and `onUpdate` when button is clicked.
 */
const PostModal: React.FC<Props> = ({ open, close, post, isUpdate, onCreate, onUpdate }) => {
    const [newTitle, setNewTitle] = useState<string>("");
    const [newDesc, setNewDesc] = useState<string>("");
    const [error, setError] = useState<string>(" ");

    const handleClose = () => {
        setNewTitle("");
        setNewDesc("");
        close();
    };

    const handleChangeTitle = (event: React.ChangeEvent<HTMLInputElement>) => {
        setNewTitle(event.target.value);
    };

    const handleChangeDesc = (event: React.ChangeEvent<HTMLInputElement>) => {
        setNewDesc(event.target.value);
    };

    /**
     * Sends the title and description to create if `isUpdate` is false or update the post
     * if `isUpdate` is true.
     */
    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setError(" ");

        if (isUpdate) {
            if (post == null) {
                setError("system error: post missing");
                return;
            }
            onUpdate(post.post_id, newTitle, newDesc);
            setNewTitle("");
            setNewDesc("");
            setError(" ");
            return;
        }
        onCreate(newTitle, newDesc);
        setNewTitle("");
        setNewDesc("");
        setError(" ");
    };

    useEffect(() => {
        if (error) setTimeout(() => setError(" "), 5000);
    }, [error]);

    useEffect(() => {
        if (isUpdate) {
            setNewTitle(post?.title || "");
            setNewDesc(post?.description || "");
        }
    }, [open, post]);

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
                {isUpdate ? "Enter the new post" : "Enter a new post."}
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
                    "& .MuiTextField-root": { mt: 2, mb: 1, mx: 3, width: "60ch" },
                }}>
                <TextField
                    required
                    id="newTitle"
                    label="New Title"
                    value={newTitle}
                    placeholder="title"
                    autoComplete="off"
                    autoFocus
                    multiline
                    rows={2}
                    onChange={handleChangeTitle}
                    error={error != " "}
                    helperText={error}
                />
                <TextField
                    required
                    id="newDesc"
                    label="New Description"
                    value={newDesc}
                    placeholder="description"
                    autoComplete="off"
                    multiline
                    rows={6}
                    onChange={handleChangeDesc}
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

export default PostModal;
