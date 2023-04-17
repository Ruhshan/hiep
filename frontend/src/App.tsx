import NavHeader from './components/NavHeader'
import React from 'react'
import {Analyzer} from './components/Analyzer';
import {ProteinFeatureViewer} from './components/ProteinFeatureViewer';

const App: React.FC = () => {
  return (
    <div className="App">
      <NavHeader></NavHeader>
        <Analyzer></Analyzer>
    </div>
  )
}

export default App
