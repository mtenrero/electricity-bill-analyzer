import React from 'react';

import './steps.css'

import axios from 'axios';

import { Upload, Icon, Button, message, Form, Spin } from 'antd';
const { Dragger } = Upload;

class Uploader extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            uploaded: false,
        }
        this.draggerProps = {
            name: 'file',
            accept: ".csv",
            multiple: false,
        };
    };

    handleUploadDone = (info) => this.setState({uploaded: true});

    normFile(e) {
        console.log('Upload event', e);
        if (Array.isArray(e)) {
            this.setState({fileUpload: e.file})
            return e;
        }
        if (e.fileList.length > 1) {
            this.setState({fileUpload: e.file})
            e.fileList.shift();
        }
        return e && e.fileList;
    }

    dummyRequest({ file, onSuccess }) {
        this.state.file = file
        const isCSV = file.name.match(/\.*.csv/i);
        if (!isCSV) {
            message.error('Debe ser un fichero con formato CSV');
        } else {
            message.success('Fichero preparado')
            onSuccess('ok');
        }
    }

    handleSubmit(event) {
        event.preventDefault();

        console.log(this.state)
        this.props.next();

        const formData = new FormData()
        formData.append('file', this.state.file)

        axios.post("http://localhost:5000/", formData, {
        }).then( res => {
            console.log(res)
            this.props.next();
        }).catch(e => {
            console.log(e);
            this.props.servererror();
        })
    }

    render() {
        const { form: { getFieldDecorator }} = this.props;
        return(
            <div className="fileDragger">
                <Spin spinning={this.props.step === 1}>
                    <Form onSubmit={this.handleSubmit.bind(this)}>
                        <Form.Item>
                            {getFieldDecorator('file', {
                                initialValie: this.props.dataset && this.props.dataset.filename ? this.props.dataset.filename : [],
                                valuePropName: 'fileList',
                                getValueFromEvent: this.normFile
                            })(
                                <Dragger {...this.draggerProps}
                                onChange={this.handleUploadDone}
                                customRequest={this.dummyRequest.bind(this)}>
                                    <p className="ant-upload-drag-icon">
                                        <Icon type="inbox" />
                                    </p>
                                    <p className="ant-upload-text">Selecciona o arrastra un fichero al área delimitada</p>
                                    <p className="ant-upload-hint">
                                        Sólo se admiten ficheros CSV
                                    </p>
                                </Dragger>
                            )}
                            
                        </Form.Item>
                        <Form.Item className="uploadConsumption">
                            <Button 
                                type="primary" 
                                shape="round" 
                                icon="cloud-upload" 
                                size="large"
                                htmlType="submit" 
                                disabled={!this.state.uploaded}
                            >
                                Analizar consumo
                            </Button>
                        </Form.Item>
                    </Form>
                </Spin>
                
            </div>
        );
    }
}

export default Form.create() (Uploader);