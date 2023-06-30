import React from 'react';
import './Table.css';

const URL = 'https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1';

const Entry = ({id, symbol, name, index}) => {
    return (
        <div className='Table-Entry'>
            <div className='Table-Entry-row'>
                {[id, symbol, name].map(i => {
                    let clazz = 'Table-Entry-col'
                    if (index === 0) clazz += ' Table-Entry-col--first'
                    if (index < 5) clazz += ' Table-Entry-col--firstof5'
                    if (symbol === 'usdt') clazz += ' Table-Entry-col--usdt'
                    return <div className={clazz}>{i}</div>
                })}
            </div>
        </div>
    );
}

const Table = () => {
    const [entries, setEntries] = React.useState([{id: 'id', symbol: 'symbol', name: 'name'}]);

    React.useEffect(() => {
        fetch(URL)
            .then(i => i.json())
            .then(i => setEntries(x => [...x, ...i]))
    }, []);

    if (!entries) return <div>Loading...</div>;
    return entries 
        ? entries.map((/** @type {object} */ i, index) => 
            <Entry {...i} index={index}/>
          )
        : <div>Loading...</div>;
}

export default Table;
