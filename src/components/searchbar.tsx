import React, { useState } from "react";
import { Box, Button, IconButton, InputAdornment, TextField } from "@mui/material";
import ClearIcon from '@mui/icons-material/Clear';
import SearchIcon from '@mui/icons-material/Search';

import type { Entity } from "../types/entity";

import axiosInstance from "../utils/axiosInstance";
import { isValidString } from "../utils/formatters";

type Props<T extends Entity> = {
    /**
     * True if T is of type Topic.
     */
    isTopic: boolean;

    /**
     * Topic id of the post. Required if `isTopic` is false.
     */
    topicId?: string;

    /**
     * Function that sets the entity list with the searched entity list.
     */
    setEntity: React.Dispatch<React.SetStateAction<T[]>>;

    /**
     * Function that sets the message.
     */
    setMessage: React.Dispatch<React.SetStateAction<string>>;

    /**
     * Function that sets the presence of error.
     */
    setIsError: React.Dispatch<React.SetStateAction<boolean>>;

    /**
     * Function that toggles SnackBar.
     */
    setOpenSnackBar: React.Dispatch<React.SetStateAction<boolean>>;
}

/**
 * Renders a search bar that sends a `GET` request to backend to search for entities that matches
 * with the query string keyed, and a `GET` request to fetch the full list when query is cleared.
 * It sets the message, error status, and toggles the SnackBar via `setMessage`, `setIsError` and
 * `setOpenSnackBar` respectively.
 */
const SearchBar = <T extends Entity,>({
    isTopic, topicId, setEntity, setMessage, setIsError, setOpenSnackBar }: Props<T>) => {
    const [query, setQuery] = useState<string>("");

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setQuery(event.target.value);
    };

    /**
     * Clears all search queries done to the entity list and sets the full list.
     */
    const handleClear = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.preventDefault();
        setOpenSnackBar(false);
        setIsError(false);
        setMessage("");

        if (isTopic) {
            axiosInstance.get("/api/topics/")
                .then(res => {
                    if (res.data) {
                        setEntity(res.data.payload?.data);
                        setIsError(false);
                        setMessage("search query cleared");
                    }
                })
                .catch(err => {
                    console.error("unable to get all topics: " + err);
                    setIsError(true);
                    setMessage(err);
                })
                .finally(() => {
                    setOpenSnackBar(true);
                    setQuery("");
                });
            return;
        }

        if (!isTopic && !topicId) {
            setIsError(true);
            setMessage("topicId misssing");
            setOpenSnackBar(true);
            return;
        }

        axiosInstance.get(`/api/posts/all/${topicId}`)
            .then(res => {
                if (res.data) {
                    setEntity(res.data.payload?.data);
                    setIsError(false);
                    setMessage("search query cleared");
                }
            })
            .catch(err => {
                console.error("unable to get all posts: " + err);
                setIsError(true);
                setMessage(err);
            })
            .finally(() => {
                setOpenSnackBar(true);
                setQuery("");
            });
    };

    /**
     * Searches the list for entities that matches the query string and sets the filtered list.
     */
    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setOpenSnackBar(false);
        setIsError(false);
        setMessage("");

        if (!isValidString(query)) {
            setMessage("Please enter a valid query");
            setIsError(true);
            setOpenSnackBar(true);
            return;
        }

        if (isTopic) {
            axiosInstance.get(`/api/topics/search?q=${query.trim()}`)
                .then(res => {
                    if (res.data) {
                        setEntity(res.data.payload?.data);
                        setIsError(false);
                        setMessage("topic list updated");
                    }
                })
                .catch(err => {
                    console.error("unable to query topics: " + err);
                    setIsError(true);
                    setMessage(err);
                })
                .finally(() => {
                    setOpenSnackBar(true);
                    setQuery("");
                });
            return;
        }

        if (!isTopic && !topicId) {
            setIsError(true);
            setMessage("topicId misssing");
            setOpenSnackBar(true);
            return;
        }

        axiosInstance.get(`/api/posts/${topicId}/search?q=${query.trim()}`)
            .then(res => {
                if (res.data) {
                    setEntity(res.data.payload?.data);
                    setIsError(false);
                    setMessage("post list updated");
                }
            })
            .catch(err => {
                console.error("unable to query posts: " + err);
                setIsError(true);
                setMessage(err);
            })
            .finally(() => {
                setOpenSnackBar(true);
                setQuery("");
            });
    };

    return (
        <Box
            component="form"
            onSubmit={handleSubmit}
            sx={{ width: "60vw", mb: 2 }}>
            <TextField
                id="search"
                value={query}
                placeholder="Search"
                autoComplete="off"
                fullWidth
                required
                onChange={handleChange}
                sx={{ "& .MuiOutlinedInput-root": { borderRadius: 10, pl: 3 } }}
                slotProps={{
                    input: {
                        startAdornment: (
                            <InputAdornment position="start">
                                <SearchIcon />
                            </InputAdornment>
                        ),
                        endAdornment: (
                            <InputAdornment position="end">
                                <IconButton
                                    onClick={handleClear}
                                    aria-label="clear"
                                    sx={{ p: "3px" }}>
                                    <ClearIcon />
                                </IconButton>
                                <Button type="submit">Search</Button>
                            </InputAdornment>
                        ),
                    }
                }}
            />
        </Box>
    )
}

export default SearchBar;
