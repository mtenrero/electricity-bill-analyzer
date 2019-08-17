import React from 'react';

import './steps.css'

import { Upload, Icon, Button } from 'antd';
const { Dragger } = Upload;

class Uploader extends React.Component {

    render() {
        return(
            <div className="fileDragger">
                <Dragger>
                    <p className="ant-upload-drag-icon">
                    <Icon type="inbox" />
                    </p>
                    <p className="ant-upload-text">Selecciona o arrastra un fichero al área delimitada</p>
                    <p className="ant-upload-hint">
                    Sólo se admiten ficheros CSV
                    </p>
                </Dragger>

                <div className="uploadConsumption">
                    <Button type="primary" shape="round" icon="cloud-upload" size="large" disabled>
                        Analizar consumo
                    </Button>
                </div>
            </div>
        );
    }
}

export default Uploader;