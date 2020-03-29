import React from 'react';
import { Card, Rate, Divider, Tooltip } from 'antd';
import {
    BrowserRouter as Router,
    Link
  } from "react-router-dom";
import '../../App.less';
import './style.less';

const { Meta } = Card;

const BusinessCard = (props) => {
    const {company} = props;
    const { Facts, Name, Rating } = company;
    return  <Card
    className="business-card"
    cover={
      <img
      style={{
        width:'100px'
      }}
        alt="example"
        src={company.Logo}
      />
    }
    actions={[
    ]}
  >
    <Link to={`/${company.Id}`}>
    <Tooltip title="See Detailed Info">
      <span className="header-sub-title">{Name} <span className="circle-rating" ><span style={{padding:'5px'}}>{Rating}</span></span></span>
      </Tooltip>
      </Link>
      {Facts && Facts.map(fact => <div>
        <Divider/>
            <span >{fact.Summary}  <a href={fact.Link}>Link</a></span>
      </div>)}
  </Card>
}

export default BusinessCard;