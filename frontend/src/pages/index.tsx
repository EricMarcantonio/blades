import React, {useCallback, useEffect, useState} from 'react'
import '@shopify/polaris/build/esm/styles.css';
import {
    AppProvider,
    Page,
    Card,
    IndexTable,
    TextStyle,
    useIndexResourceState,
    Modal,
    TextContainer,
    Button,
    Form,
    FormLayout,
    Checkbox,
    TextField
} from '@shopify/polaris';
import axios from "axios";
import moment from "moment";

interface IProduct {
    name: string
    price: number
    id: number
    units: number
    modified_date: string
    added_date: string
}


const url = process.env.GATSBY_DOMAIN
const port = process.env.GATSBY_PORT

export const post = async (body: string) => {
    return await axios.post(`${url}:${port}/gql`, {
        query: body
    }).then((r) => {
        return r.data
    })
}


const HelloWorld = () => {
    let temp: IProduct = {name: "", price: 0, units: 0, id: 0, modified_date: "", added_date: ""}
    const [rows, setRows] = useState<IProduct[]>(Array(temp))

    // @ts-ignore
    const {selectedResources, allResourcesSelected, handleSelectionChange} = useIndexResourceState(rows);
    const handleNameChange = useCallback((value) => setName(value), []);
    const handlePriceChange = useCallback((value) => setPrice(value), []);
    const handleUnitChange = useCallback((value) => setUnits(value), []);
    const [active, setActive] = useState(false);
    const handleChange = useCallback(() => {
        setActive(!active)
    }, [active]);
    const [id, setId] = useState("0")
    const [name, setName] = useState('');
    const [price, setPrice] = useState(0);
    const [units, setUnits] = useState(0)




    const updateRows = async () => {
        await post(`query {products { id, name, added_date, modified_date, price, units } }`).then(r => {
            setRows(r.data.products);
        })
    }


    const resourceName = {
        singular: 'skate', plural: 'skates',
    };

    const HandleBulkDelete = async () => {
        if (allResourcesSelected) {
            await Promise.all(rows.map(v => {
                return post(`mutation { deactivateProduct(id: ${v.id}){ id }}`)
            })).then(async r => {
                await updateRows()
            })
        }
        await Promise.all(selectedResources.map(v => {
            return post(`mutation { deactivateProduct(id: ${v}){ id }}`)
        })).then(async r => {
            await updateRows();
        })
    }


    const promotedBulkActions = [{
        content: 'Edit Skate', onAction: () => {

            setName(rows.filter(r => r.id.toString() == selectedResources[0])[0] ? rows.filter(r => r.id.toString() == selectedResources[0])[0].name : "")
            setPrice(rows.filter(r => r.id.toString() == selectedResources[0])[0] ? rows.filter(r => r.id.toString() == selectedResources[0])[0].price : 0)
            setUnits(rows.filter(r => r.id.toString() == selectedResources[0])[0] ? rows.filter(r => r.id.toString() == selectedResources[0])[0].units : 0)
            handleChange()
        },
    }, {
        content: 'Delete Skate', onAction: HandleBulkDelete,
    },];


    const handleSubmit = useCallback((_event) => {
        setName('');
    }, []);


    const seedDB = async () => {
        await axios.get(`${url}:${port}/?seed=yes`)
        updateRows()
    }

    const updateProduct = async () => {
        let queryString: string
        if (allResourcesSelected){
            queryString = `mutation { updateProduct(id: ${rows[0].id}, price: ${price}, units: ${units}, name: "${name}" ) { id }}`
        } else {
            queryString = `mutation { updateProduct(id: ${selectedResources[0]}, price: ${price}, units: ${units}, name: "${name}" ) { id }}`
        }
        await post(queryString).then(() => {
            updateRows()
            handleChange()
        }).catch(e => {
            console.log(e)
        })
    }

    const rowMarkup = () => {
        if (rows) {
            return rows.map(({id, name, added_date, modified_date, price, units}, index) => (<IndexTable.Row
                id={id.toString()}
                key={id.toString()}
                selected={selectedResources.includes(id.toString())}
                position={index}
            >
                <IndexTable.Cell>
                    <TextStyle variation="strong">{id}</TextStyle>
                </IndexTable.Cell>
                <IndexTable.Cell>{name}</IndexTable.Cell>
                <IndexTable.Cell>{price}</IndexTable.Cell>
                <IndexTable.Cell>{units}</IndexTable.Cell>
                <IndexTable.Cell>{moment(modified_date).format("MMM Do YY, h:mm:ss a")}</IndexTable.Cell>
                <IndexTable.Cell>{moment(added_date).format("MMM Do YY, h:mm:ss a")}</IndexTable.Cell>

            </IndexTable.Row>),)
        } else {
            return <div>Loading</div>
        }
    }

    useEffect(() => {
        updateRows()
    }, [])

    return (<AppProvider i18n={{}}>
        <Page title="BLADES Inventory">
            <Card
                primaryFooterAction={{
                    async onAction() {
                        await seedDB()

                    }, destructive: true, content: "Seed the DB"
                }}>
                <IndexTable
                    resourceName={resourceName}
                    itemCount={rows ? rows.length : 0}
                    selectedItemsCount={allResourcesSelected ? 'All' : selectedResources.length}
                    onSelectionChange={handleSelectionChange}
                    promotedBulkActions={promotedBulkActions}
                    bulkActions={promotedBulkActions}
                    headings={[{title: 'SKU'}, {title: 'Name'}, {title: 'Price'}, {title: 'Units'}, {title: 'Last Modified'}, {title: 'Added On'}]}
                >
                    {rowMarkup()}
                </IndexTable>
            </Card>


            <div style={{height: '500px'}}>
                <Modal
                    open={active}
                    onClose={handleChange}
                    title={`Edit ${name}`}
                    primaryAction={{
                        content: 'Submit', onAction: async () => await updateProduct(),
                    }}
                    secondaryActions={[{
                        content: 'Cancel', onAction: handleChange,
                    },]}

                >
                    <Modal.Section>
                        <TextContainer>
                            <Form onSubmit={updateProduct}>
                                <FormLayout>
                                    <TextField
                                        value={name}
                                        onChange={handleNameChange}
                                        label="Name"
                                        autoComplete="off"
                                        helpText={<span>The name of the skate</span>}
                                    />
                                    <TextField
                                        value={price.toString()}
                                        onChange={handlePriceChange}
                                        label="Price"
                                        autoComplete="off"
                                        helpText={<span>The price of the skate</span>}
                                    />
                                    <TextField
                                        value={units.toString()}
                                        onChange={handleUnitChange}
                                        label="Units"
                                        autoComplete="off"
                                        helpText={<span>The number of skates of this type</span>}
                                    />
                                </FormLayout>
                            </Form>
                        </TextContainer>
                    </Modal.Section>
                </Modal>
            </div>
        </Page>
    </AppProvider>)
}

export default HelloWorld;
