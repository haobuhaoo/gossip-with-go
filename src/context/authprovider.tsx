import { useEffect, useState } from "react";

import type { User } from "../types/entity";

import { AuthContext } from "./authcontext";

import axiosInstance from "../utils/axiosInstance";

/**
 * Wraps the application and provides authentication state. It verifies the user via `GET` request.
 */
export default function AuthProvider({ children }: { children: React.ReactNode }) {
    const [auth, setAuth] = useState<User>({
        username: "",
        userId: "",
        isAuthenticated: false,
        isLoading: true,
    })

    useEffect(() => {
        const token: string | null = localStorage.getItem("token");
        if (!token) {
            setAuth({ username: "", userId: "", isAuthenticated: false, isLoading: false });
            return;
        }
        axiosInstance.get("/api/me")
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
                localStorage.removeItem("token");
                setAuth({ username: "", userId: "", isAuthenticated: false, isLoading: false });
            });
    }, []);

    return (
        <AuthContext.Provider value={{ auth, setAuth }}>
            {children}
        </AuthContext.Provider>
    )
}
