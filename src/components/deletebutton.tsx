import React from "react";
import { IconButton } from "@mui/material";
import DeleteIcon from "@mui/icons-material/Delete";

import type { Entity } from "../types/entity";

type Props<T extends Entity> = {
    /**
     * The entity of interest.
     */
    entity: T;

    /**
     * Function that passes the entity to be deleted back to parent component.
     */
    onDelete: (e: T) => void;
};

/**
 * Renders a delete button which calls `onDelete` when clicked.
 */
const DeleteButton = <T extends Entity,>({ entity, onDelete }: Props<T>) => {
    const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.stopPropagation();
        onDelete(entity);
    };

    return (
        <IconButton
            onClick={handleClick}
            aria-label="delete"
            sx={{ p: "3px", borderRadius: 10, color: "red" }}>
            <DeleteIcon />
        </IconButton>
    )
}

export default DeleteButton;
