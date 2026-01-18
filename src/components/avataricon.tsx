import React from "react";
import { Avatar } from "@mui/material";

import { showInitial } from "../utils/formatters";

type Props = {
    /**
     * Username of the author.
     */
    username: string;
}

/**
 * Renders an avatar icon of the author.
 */
const AvatarIcon: React.FC<Props> = ({ username }) => {
    return <Avatar>{showInitial(username)}</Avatar>
}

export default AvatarIcon;
