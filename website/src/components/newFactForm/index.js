import React from 'react';
import { Form, Input, Select, Tooltip, Button } from 'antd';
import './style.less'
const {TextArea } = Input;
const NewFactForm = () => {
    const onFinish = values => {
        console.log('Received values of form: ', values);
      };
      return (
          <div className="new-fact-form">
  <Form name="complex-form" onFinish={onFinish} labelCol={{ span: 8 }} wrapperCol={{ span: 16 }}>
          <Form.Item label="Citation">
            <Form.Item
              name="citation"
              noStyle
              rules={[{ required: true, message: 'citation is required' }]}
            >
              <Input style={{ width: 160 }} placeholder="Please enter citation" />
            </Form.Item>
          </Form.Item>

          <Form.Item label="Summary">
            <Form.Item
              name="summary"
              noStyle
              rules={[{ required: true, message: 'summary is required' }]}
            >
              <TextArea style={{ width: 300 }} placeholder="Please enter summary" />
            </Form.Item>
          </Form.Item>
          
          <Form.Item label=" " colon={false}>
            <Button type="primary" htmlType="submit">
              Submit
            </Button>
            <Button type="primary">
              Cancel
            </Button>
          </Form.Item>
        </Form>
          </div>
      )
}

export default NewFactForm;