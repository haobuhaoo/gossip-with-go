import React from "react";
import { IconButton } from "@mui/material";
import EditIcon from "@mui/icons-material/Edit";

import type { Entity } from "../types/entity";

type Props<T extends Entity> = {
    /**
     * The entity of interest.
     */
    entity: T;

    /**
     * Function that passes the entity to be updated back to parent component.
     */
    updateEntity: (e: T) => void;
};

/**
 * Renders an edit button which calls `updateEntity` when clicked.
 */
const EditButton = <T extends Entity,>({ entity, updateEntity }: Props<T>) => {
    const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.stopPropagation();
        updateEntity(entity);
    };

    return (
        <IconButton
            onClick={handleClick}
            aria-label="edit"
            sx={{
                p: "3px",
                borderRadius: 10,
                color: "blue",
                "&:hover": { backgroundColor: "lightgrey" }
            }}>
            <EditIcon />
        </IconButton>
    )
}

export default EditButton;
