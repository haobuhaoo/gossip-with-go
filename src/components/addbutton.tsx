import React from "react";
import { Button } from "@mui/material";
import AddIcon from '@mui/icons-material/Add';

type Props = {
    /**
     * Function that sets to open or close the modal.
     */
    setOpenModal: (b: boolean) => void;
}

/**
 * Renders an add button which calls `setOpenModal` when clicked.
 */
const AddButton: React.FC<Props> = ({ setOpenModal }) => {
    return (
        <Button
            variant="outlined"
            size="large"
            onClick={() => setOpenModal(true)}
            sx={{
                position: "absolute",
                top: 48,
                right: 80,
                borderRadius: 3,
                fontSize: "20px",
                "&:hover": { backgroundColor: "#5aacfdff", color: "white" }
            }}>
            <AddIcon sx={{ display: "flex", fontSize: "24px", mr: 0.25 }} />
            Add
        </Button>
    )
}

export default AddButton;
