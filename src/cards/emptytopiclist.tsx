import React from "react";
import { Typography } from "@mui/material";
import NoteAddOutlinedIcon from '@mui/icons-material/NoteAddOutlined';

/**
 * Renders an empty topic notice.
 */
const EmptyTopicList: React.FC = () => {
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
                No topics available.
            </Typography>
        </div>
    )
}

export default EmptyTopicList;
