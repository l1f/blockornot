import { Typography, useTheme } from "@mui/material";
import React from "react";
import { NavLink } from "react-router-dom";

const Footer = () => {
  const theme = useTheme();
  return (
    <Typography component="footer" sx={{ backgroundColor: "primary.main", width: "100%", padding: 1, textAlign: "center" }}>
      <NavLink to="/about" style={{ color: theme.palette.primary.contrastText }}>About</NavLink>
    </Typography>
  );
};

export default Footer;
