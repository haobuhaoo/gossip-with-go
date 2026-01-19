import { useEffect, useState } from "react";

import type { User } from "../types/entity";

import { AuthContext } from "./authcontext";

import axiosInstance from "../utils/axiosInstance";

/**
 * Wraps the application and provides authentication state. It verifies the user via `GET` request.
 */
export default function AuthProvider({ children }: { children: React.ReactNode }) {
    const [auth, setAuth] = useState<User>({
        username: localStorage.getItem("user") ?? "",
        userId: localStorage.getItem("token") ?? "",
        isAuthenticated: false,
        isLoading: true,
    })

    useEffect(() => {
        if (!auth.username && !auth.userId) return;

        axiosInstance.get(`/users/${auth.username}`)
            .then(res => {
                if (res.data) {
                    setAuth({
                        username: res.data.payload?.data?.name,
                        userId: res.data.payload?.data?.user_id,
                        isAuthenticated: true,
                        isLoading: false,
                    });
                }
            })
            .catch(err => {
                console.error("unable to verify user: " + err);
                localStorage.removeItem("user");
                localStorage.removeItem("token");
                setAuth({ username: "", userId: "", isAuthenticated: false, isLoading: false });
            });
    }, [auth.username, auth.userId]);

    return (
        <AuthContext.Provider value={{ auth, setAuth }}>
            {children}
        </AuthContext.Provider>
    )
}
