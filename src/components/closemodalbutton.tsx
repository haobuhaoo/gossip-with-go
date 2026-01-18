import React from "react";
import { IconButton } from "@mui/material";
import CloseIcon from '@mui/icons-material/Close';

type Props = {
    /**
     * Function that closes the modal.
     */
    close: () => void;
}

/**
 * Renders a close button that calls `close` when clicked.
 */
const CloseModalButton: React.FC<Props> = ({ close }) => {
    return (
        <IconButton
            onClick={close}
            sx={{
                position: "absolute",
                top: 24,
                right: 20,
                fontSize: "24px",
                borderRadius: 10,
                cursor: "pointer",
                "&:hover": { backgroundColor: "lightgray", }
            }}>
            <CloseIcon />
        </IconButton>
    )
}

export default CloseModalButton;
