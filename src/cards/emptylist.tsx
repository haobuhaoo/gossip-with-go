import React from "react";
import { Typography } from "@mui/material";
import NoteAddOutlinedIcon from '@mui/icons-material/NoteAddOutlined';

type Props = {
    /**
     * The entity string to be displayed.
     */
    entity: string;
}

/**
 * Renders an empty entity notice.
 */
const EmptyList: React.FC<Props> = ({ entity }) => {
    return (
        <div
            style={{
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
                flexDirection: "column"
            }}>
            <NoteAddOutlinedIcon color="action" sx={{ height: "128px", width: "128px" }} />

            <Typography variant="h6" gutterBottom sx={{ fontWeight: "bold" }}>
                No {entity} available.
            </Typography>
        </div>
    )
}

export default EmptyList;
