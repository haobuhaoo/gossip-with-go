import React from "react";
import { IconButton, Typography } from "@mui/material";
import ThumbDownIcon from '@mui/icons-material/ThumbDown';
import ThumbUpIcon from '@mui/icons-material/ThumbUp';

type Props = {
    type: "like" | "dislike";

    vote: 1 | -1 | null;

    voteCount: number;
}

const VoteButton: React.FC<Props> = ({ type, vote, voteCount }) => {
    return (
        <IconButton sx={{ gap: 1, borderRadius: 2 }}>
            {type == "like"
                ? <ThumbUpIcon sx={{ color: `${vote == 1 ? "red" : ""}` }} />
                : <ThumbDownIcon sx={{ color: `${vote == -1 ? "red" : ""}` }} />
            }

            <Typography sx={{ fontSize: "18px", color: "black" }}>
                {voteCount}
            </Typography>
        </IconButton>
    )
}

export default VoteButton;
