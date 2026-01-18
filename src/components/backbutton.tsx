import React from "react";
import { Button } from "@mui/material";
import ArrowBackIosIcon from '@mui/icons-material/ArrowBackIos';

type Props = {
    /**
     * Function that navigates back to the previous page.
     */
    handleBack: () => void;
}

/**
 * Renders a back button which calls `handleBack` when clicked.
 */
const BackButton: React.FC<Props> = ({ handleBack }) => {
    return (
        <Button
            variant="outlined"
            size="large"
            onClick={handleBack}
            sx={{
                position: "absolute",
                top: 48,
                left: 80,
                borderRadius: 3,
                fontSize: "20px",
                "&:hover": { backgroundColor: "#5aacfdff", color: "white" }
            }}>
            <ArrowBackIosIcon sx={{ fontSize: "20px" }} />
            Back
        </Button>
    )
}

export default BackButton;
