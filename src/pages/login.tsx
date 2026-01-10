import React, { useEffect, useRef, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Box, Button, Card, TextField, Typography } from "@mui/material";

import axiosInstance from "../utils/axiosInstance";

/**
 * Models a login form page that asks for a username and sends a `GET` or `POST` request to fetch or
 * create user data. On success, the username is stored as a token in `localStorage`.
 */
const LoginPage: React.FC = () => {
    const [error, setError] = useState<string>(" ");
    const [username, setUsername] = useState<string>("");
    const [isLogin, setIsLogin] = useState<boolean>(true);
    const navigate = useNavigate();
    const usernameRef = useRef<HTMLInputElement | null>(null);

    const handleButtonClick = () => {
        setIsLogin(!isLogin);
        setUsername("");
        setError(" ");
        setTimeout(() => usernameRef.current?.focus(), 600);
    };

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setUsername(event.target.value);
    };

    /**
     * Sends the username to backend API to fetch user data or create a new user. On success, the
     * username is stored as a token and the user is navigated to the home page.
     */
    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setError(" ");

        if (isLogin) {
            axiosInstance.get(`users/${username}`)
                .then(res => {
                    if (res.data) {
                        localStorage.setItem("user", res.data.payload?.data?.name);
                        localStorage.setItem("token", res.data.payload?.data?.user_id);
                        navigate("/home");
                    }
                })
                .catch(err => {
                    console.error("login error: " + err);
                    setError(err);
                });
        } else {
            axiosInstance.post("/users", { name: username })
                .then(res => {
                    if (res.data) {
                        localStorage.setItem("user", res.data.payload?.data?.name);
                        localStorage.setItem("token", res.data.payload?.data?.user_id);
                        navigate("/home");
                    }
                })
                .catch(err => {
                    console.error("register error: " + err)
                    setError(err);
                });
        }
    };

    useEffect(() => {
        if (error) {
            setTimeout(() => setError(" "), 5000);
        }
    }, [error]);

    return (
        <div
            style={{
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
                height: "100vh",
                width: "100vw"
            }}>
            <Card
                variant="outlined"
                sx={{
                    display: "flex",
                    width: "60vw",
                    height: "60vh",
                    borderRadius: "20px",
                    boxShadow: "0 14px 28px rgba(0, 0, 0, 0.25), 0 10px 10px rgba(0, 0, 0, 0.22)"
                }}>
                <Box
                    sx={{
                        display: "flex",
                        justifyContent: "center",
                        alignItems: "center",
                        flexDirection: "column",
                        width: "45%",
                        backgroundColor: "#5aacfdff",
                        transition: "transform 0.6s ease",
                        transform: isLogin ? "translateX(0%)" : "translateX(122.22%)",
                        zIndex: 2,
                    }}>
                    <Typography variant="h4" sx={{ color: "#FFF", fontWeight: "bold" }}>
                        {isLogin ? "Hello, Welcome!" : "Welcome Back!"}
                    </Typography>
                    <Typography
                        variant="caption"
                        sx={{
                            color: "#FFF",
                            fontSize: "16px",
                            paddingY: 1

                        }}>
                        {isLogin ? "Don't have an account?" : "Already have an account?"}
                    </Typography>
                    <Button
                        variant="outlined"
                        onClick={handleButtonClick}
                        sx={{
                            border: "1px solid white",
                            color: "white",
                            fontWeight: 600,
                            "&:hover": { backgroundColor: "#4c94dbff" }
                        }}>
                        {isLogin ? "Register" : "Login"}
                    </Button>
                </Box>

                <Box
                    component="form"
                    onSubmit={handleSubmit}
                    sx={{
                        display: "flex",
                        justifyContent: "center",
                        alignItems: "center",
                        width: "55%",
                        flexDirection: "column",
                        "& .MuiTextField-root": { mt: 2, mb: 1, mx: 3, width: "40ch" },
                        transition: "transform 0.6s ease",
                        transform: isLogin ? "translateX(0%)" : "translateX(-81.82%)",
                    }}>
                    <Typography
                        variant="h4"
                        sx={{
                            display: "flex",
                            justifyContent: "center",
                            alignItems: "center",
                            fontWeight: "bold",
                            marginBottom: 1
                        }}>
                        {isLogin ? "Login" : "Registeration"}
                    </Typography>
                    <TextField
                        required
                        id="username"
                        label="Username"
                        name="username"
                        value={username}
                        placeholder="username"
                        autoComplete="off"
                        autoFocus
                        inputRef={usernameRef}
                        onChange={handleChange}
                        error={error != " "}
                        helperText={error}
                    />
                    <Button
                        variant="contained"
                        type="submit"
                        sx={{
                            display: "flex",
                            justifyContent: "center",
                            alignItems: "center",
                            height: "5ch",
                            width: "20ch",
                            fontSize: "16px"
                        }}>
                        {isLogin ? "Login" : "Register"}
                    </Button>
                </Box>
            </Card>
        </div>
    );
};

export default LoginPage;
