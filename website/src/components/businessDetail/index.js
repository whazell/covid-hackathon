import React, { useState } from 'react';
import {
    useParams
  } from "react-router-dom";
  import { List } from 'antd';

import '../../App.less';
import './style.less';
import { useEffectAsync } from '../../util';
import { getCompanyDetail } from '../../api/company';
import { FactForm } from '../index';
import { Button } from 'antd';
export const BusinessDetail = () => {
    let { id } = useParams();
    const [business, setBusiness ] = useState({});
    const [showFactForm, setShowFactForm] = useState(false);
    useEffectAsync(async () => {
    const res = await getCompanyDetail(id);
    setBusiness(res.data);
    }, [id])
    return (
        <div className="business-detail">
      <p className="business-detail-header-text">
           {business.Name}
        </p>
        <img style={{width:'150px'}} src={business.Logo}/>
        <Button style={{paddingLeft:'20px', marginLeft:'20px'}} onClick={() => {
      setShowFactForm(true)
    }}>Submit New Fact</Button>
        <List
    itemLayout="horizontal"
    dataSource={business.Facts}
    renderItem={item => (
      <List.Item>
        <List.Item.Meta
          title={<a href={item.Citation}>Link</a>}
          description={item.Summary}
        />
      </List.Item>
    )}
  />
  <FactForm visible={showFactForm} id= {id} onCancel={() => {
      setShowFactForm(false)
  }}/>
        </div>
    )
}

export default BusinessDetail;