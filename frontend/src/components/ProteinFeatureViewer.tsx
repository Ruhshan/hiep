import {FeatureViewer} from 'feature-viewer-typescript/src/feature-viewer';
import React, {useEffect, useState} from 'react';
import ApiResult from '../domains/ApiResult';

interface Props {
    apiResult: ApiResult;
}


export function ProteinFeatureViewer() {
    const initViewer = (divId: string) => {
        const P04050 =
            'MTKFTILLISLLFCIAHTCSASKWQHQQDSCRKQLQGVNLTPCEKHIMEKIQGRGDDDDDDDDDNHILRTMRGRINYIRRNEGKDEDEEEEGHMQKCCTEMSELRSPLMTKFTILLISLLFCIAHTCSASKWQHQQDSCRKQLQGVNLTPCEKHIMEKIQGRGDDDDDDDDDNHILRTMRGRINYIRRNEGKDEDEEEEGHMQKCCTEMSELRSPLMTKFTILLISLLFCIAHTCSASKWQHQQDSCRKQLQGVNLTPCEKHIMEKIQGRGDDDDDDDDDNHILRTMRGRINYIRRNEGKDEDEEEEGHMQKCCTEMSELRSPL';
        const fv = new FeatureViewer(P04050, '#' + divId, {
            toolbar: false,
            toolbarPosition: 'left',
            brushActive: true,
            zoomMax: 5,
            flagColor: '#DFD5F5',
            animation: true,
            showSequence: true,
        });

        return fv;
    };

    const viewViewer = (divId: string) => {
        document.getElementById(divId).innerHTML = ''

        const fv = initViewer(divId);
        const featurelist = [
            {
                type: 'rect',
                id: 'mysimplefeature',
                data: [{ x: 50, y: 100 }],
            },
            {
                type: 'rect',
                id: 'mysimplefeature2',
                data: [{ x: 70, y: 90 }],
            },
        ];
         fv.addFeatures(featurelist)
    };

    useEffect(() => {


        viewViewer('fvDivInit');




    }, []);

    return <div id="fvDivInit" />;
}
