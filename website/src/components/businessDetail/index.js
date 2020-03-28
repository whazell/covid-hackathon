import React from 'react';
import {
    useParams
  } from "react-router-dom";
export const BusinessDetail = () => {
    let { id } = useParams();
    return (
        <p>
            This is business detail page {id}
        </p>
    )
}

export default BusinessDetail;