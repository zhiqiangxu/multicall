// SPDX-License-Identifier: MIT
pragma solidity >=0.5.15;
pragma experimental ABIEncoderV2;

contract MultiCall {
    constructor(address[] memory targets, bytes[] memory datas) {
        uint256 len = targets.length;
        require(datas.length == len, "Error: Array lengths do not match.");

        bytes[] memory returnDatas = new bytes[](len);

        for (uint256 i = 0; i < len; i++) {
            address target = targets[i];
            bytes memory data = datas[i];
            (bool success, bytes memory returnData) = target.call(data);
            if (!success) {
                if (returnData.length == 0) {
                    revert("Error: call failed with no reason.");
                } else {
                    assembly {
                        let returndata_size := mload(returnData)
                        revert(add(32, returnData), returndata_size)
                    }
                }
            }
            returnDatas[i] = returnData;
        }
        bytes memory result = abi.encode(block.number, returnDatas);
        assembly {
            return(add(result, 32), mload(result))
        }
    }
}
