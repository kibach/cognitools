export const getAttribute = (attributeList, attrName) => ((attributeList || []).find((attr) => attr.Name === attrName) || {}).Value;
export const getUpdatableAttributes = (attributeList) => (attributeList || []).filter((attr) => attr.Name !== 'sub');
