import React from 'react';
import { Form, Input, Select, Tooltip, Button, message, Modal} from 'antd';
import { postCompanyFact } from '../../api/company';
const {TextArea } = Input;
const FactForm = (props) => {
  const { visible, onCancel,id} = props;
    const onFinish = async values => {
       const { condition, summary } = values;
       try {
        await postCompanyFact(id, condition, summary);
        message.success('Post new fact successfully');
       }
       catch(e) {
         
       }
      };
      return (
          <div className="new-fact-form">
            <Modal visible={visible} 
            footer={null}
            onCancel={() => onCancel()}
            title="Submit New Fact"
            >
            <Form name="complex-form" onFinish={onFinish} labelCol={{ span: 4 }} wrapperCol={{ span: 16 }}>
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
              <TextArea style={{ width: 250 }} placeholder="Please enter summary" />
            </Form.Item>
          </Form.Item>
          
          <Form.Item label=" " colon={false}>
            <Button style={{marginRight: '20px', width: '120px'}} type="primary" htmlType="submit">
              Submit
            </Button>
            <Button onClick={() => onCancel()}style={{marginRight: '20px', width: '120px'}}>
              Cancel
            </Button>
          </Form.Item>
        </Form>
            </Modal>
          </div>
      )
}

export default FactForm;