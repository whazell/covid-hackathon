import React from 'react';
import { Card, Rate, Divider } from 'antd';
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link
  } from "react-router-dom";
import '../../App.less';
import './style.less';

const { Meta } = Card;

const BusinessCard = (props) => {
    const {company} = props;
    const { Facts, Name, Rating } = company;
    console.log(Facts);
    return  <Card
    className="business-card"
    cover={
      <img
        alt="example"
        src={company.Logo}
      />
    }
    actions={[
    ]}
  >
    <Link to="/netflix">
      <span className="header-sub-title">{Name} <span className="circle-rating" ><span style={{padding:'5px'}}>{Rating}</span></span></span>
      </Link>
      {Facts.map(fact => <div>

        <Divider/>
            <span >{fact.Summary}  <a href={fact.Link}>Link</a></span>
      </div>)}
  </Card>
}

export default BusinessCard;