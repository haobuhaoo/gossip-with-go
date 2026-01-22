import React from "react";
import { IconButton, Typography } from "@mui/material";
import ThumbDownIcon from '@mui/icons-material/ThumbDown';
import ThumbUpIcon from '@mui/icons-material/ThumbUp';

type Props = {
    /**
     * The type of vote button.
     */
    type: "like" | "dislike";

    /**
     * Values the vote can hold. 1 for liked, -1 for disliked, null for not voted.
     */
    vote: 1 | -1 | null;

    /**
     * Total number of votes for the entity.
     */
    voteCount: number;

    /**
     * The entity id number.
     */
    id: number;

    /**
     * The type of entity.
     */
    entityType: "comment" | "post";

    /**
     * Function that passes the entity id number back to the parent component to indicate a like
     * for the entity.
     */
    onLike: (id: number, entity: string) => void;

    /**
     * Function that passes the entity id number back to the parent component to indicate a dislike
     * for the entity.
     */
    onDislike: (id: number, entity: string) => void;

    /**
     * Function that passes the entity id number back to the parent component to remove user's vote
     * for the entity.
     */
    onRemoveVote: (id: number, entity: string) => void;
}

/**
 * Renders a like or dislike button. It calls `onLike` or `onDislike` depending on the type of the
 * button, and `onRemoveVote` when clicked again after voting.
 */
const VoteButton: React.FC<Props> = ({
    type, vote, voteCount, id, entityType, onLike, onDislike, onRemoveVote }) => {
    const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.stopPropagation();
        if (type == "like") {
            if (vote == 1) {
                onRemoveVote(id, entityType);
            }
            onLike(id, entityType);
        } else if (type == "dislike") {
            if (vote == -1) {
                onRemoveVote(id, entityType);
            }
            onDislike(id, entityType);
        }
    };

    return (
        <IconButton onClick={handleClick} aria-label="vote-button" sx={{ gap: 1, borderRadius: 2 }}>
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
