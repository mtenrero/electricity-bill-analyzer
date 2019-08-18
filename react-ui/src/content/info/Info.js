import React from 'react';
import { Alert } from 'antd';

function Info() {
    return(
        <div>
            <Alert
                message="¿Cómo funciona el analizador?"
                description="Necesitas disponer del fichero CSV que te proporciona tu compañía eléctrica, 
                después, debes subirlo desde el formulario que aparece en la sección inferior y pulsar enviar.  
                El programa emitirá el informe de gasto que has realizado"
                type="info"
                showIcon
            />
            <br></br>
            <Alert message="Los datos se envían a nuestros servidores, pero son desechados tras su análisis. Cumplimos la RGPD / LOPD" type="warning" showIcon />
        </div>
    )
}

export default Info;