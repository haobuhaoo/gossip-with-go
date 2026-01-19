import React from "react";
import { Navigate, Outlet } from "react-router-dom";
import { CircularProgress } from "@mui/material";

import { useAuth } from "./authcontext";

/**
 * Allows access to nested routes only to authenticated users. Unauthenticated users are
 * redirected to the login page.
 */
const ProtectedRoute: React.FC = () => {
    const { auth } = useAuth();

    if (auth.isLoading) {
        return <CircularProgress />;
    }

    return (
        auth.isAuthenticated
            ? <Outlet />
            : <Navigate to="/" replace state={{ message: "Please login" }} />
    )
}

export default ProtectedRoute;
