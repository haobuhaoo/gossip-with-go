import { createContext, useContext } from "react";

import type { User } from "../types/entity";

type AuthContextType = {
    /**
     * Current authentication state.
     */
    auth: User;

    /**
     * Function to update the authentication state.
     */
    setAuth: React.Dispatch<React.SetStateAction<User>>;
}

export const AuthContext = createContext<AuthContextType | undefined>(undefined);

/**
 * Allows use of authentication context.
 * @returns Authentication context value.
 * @throws Error if used outside of AuthProvider.
 */
export function useAuth(): AuthContextType {
    const context: AuthContextType | undefined = useContext(AuthContext);

    if (!context) throw new Error("useAuth must be used within an AuthProvider");

    return context;
}
