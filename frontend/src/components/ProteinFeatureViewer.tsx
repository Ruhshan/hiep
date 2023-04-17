import {FeatureViewer} from 'feature-viewer-typescript/src/feature-viewer';
import React, {useEffect, useState} from 'react';
import ApiResult from '../domains/ApiResult';
import {FeaturesList} from 'feature-viewer-typescript/src/interfaces';
import {Box, Center} from '@chakra-ui/react';

interface Props {
    apiResult: ApiResult;
}


export function ProteinFeatureViewer(props: Props) {
    const apiResult: ApiResult = props.apiResult;
    const initViewer = (divId: string) => {

        const fv = new FeatureViewer(apiResult.querySequence, '#' + divId, {
            toolbar: false,
            toolbarPosition: 'left',
            brushActive: true,
            zoomMax: 5,
            flagColor: '#FC8181',
            animation: true,
            showSequence: true
        });

        return fv;
    };

    const styles = {
        fvBox: {
            width: '100%',
            border: '1px solid rgba(0, 0, 0, 0.36)',
            borderRadius: '5px'
        }
    }

    const viewViewer = (divId: string) => {
        document.getElementById(divId).innerHTML = '';

        const fv = initViewer(divId);

        const featurelist = apiResult.sequenceAndPositions.map(snp => ({
            type: 'rect',
            label: `${snp.position[0]} to ${snp.position[1]} aa`,
            id: `f${snp.position[0]}To${snp.position[1]}Aa`,
            color: '#22c35e',
            data: [{x: 1+snp.position[0], y: snp.position[1]}]
        })) as FeaturesList;

        console.log(featurelist);

        fv.addFeatures(featurelist);
    };

    useEffect(() => {
        if (apiResult.querySequence) {
            viewViewer('fvDivInit');
        } else {
            document.getElementById('fvDivInit').innerHTML = '';
        }


    }, [props.apiResult]);

    return <Center width={'100%'}>

                <div style={styles.fvBox}>
                    <div id="fvDivInit"/>
                </div>

            </Center>;
}
