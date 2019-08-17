import React from 'react';

import './steps.css'

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

    beforeUpload(file) {
        const isCSV = file.name.match(/\.*.csv/i);
        if (!isCSV) {
            let fileList = [];
            this.setState({fileList})
            message.error('Debe ser un fichero con formato CSV');
        } else {
            let fileList = [file];
            this.setState({fileList})
            message.success('Fichero preparado')
            return false
        }
    }

    handleUploadDone = (info) => this.setState({uploaded: true});

    normFile(e) {
        console.log('Upload event', e);
        if (Array.isArray(e)) {
            return e;
        }
        if (e.fileList.length > 1) {
            e.fileList.shift();
        }
        return e && e.fileList;
    }

    dummyRequest({ file, onSuccess }) {
        const isCSV = file.name.match(/\.*.csv/i);
        if (!isCSV) {
            message.error('Debe ser un fichero con formato CSV');
        } else {
            message.success('Fichero preparado')
            onSuccess('ok');
        }
    }

    render() {
        const { form: { getFieldDecorator }} = this.props;
        return(
            <div className="fileDragger">
                <Spin spinning={this.props.step === 1}>
                    <Form.Item>
                        {getFieldDecorator('file', {
                            initialValie: this.props.dataset && this.props.dataset.filename ? this.props.dataset.filename : [],
                            valuePropName: 'fileList',
                            getValueFromEvent: this.normFile
                        })(
                            <Dragger {...this.draggerProps}
                            onChange={this.handleUploadDone}
                            customRequest={this.dummyRequest}>
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
                </Spin>
                <div className="uploadConsumption">
                    <Button 
                        type="primary" 
                        shape="round" 
                        icon="cloud-upload" 
                        size="large" 
                        disabled={!this.state.uploaded}
                        onClick={this.props.next}
                    >
                        Analizar consumo
                    </Button>
                </div>
            </div>
        );
    }
}

export default Uploader;