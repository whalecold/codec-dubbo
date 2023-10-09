/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package org.apache.dubbo.tests.provider;

import org.apache.dubbo.tests.api.*;

import java.util.Map;
import java.util.List;

public class UserProviderImpl implements UserProvider {
    @Override
    public Boolean EchoBool(Boolean req) throws Exception {
        return req;
    }

    @Override
    public Byte EchoByte(Byte req) throws Exception {
        return req;
    }

    @Override
    public Short EchoInt16(Short req) throws Exception {
        return req;
    }

    @Override
    public Integer EchoInt32(Integer req) throws Exception {
        return req;
    }

    @Override
    public Long EchoInt64(Long req) throws Exception {
        return req;
    }

    @Override
    public Double EchoDouble(Double req) throws Exception {
        return req;
    }

    @Override
    public String EchoString(String req) throws Exception {
        return req;
    }

    @Override
    public byte[] EchoBinary(byte[] req) throws Exception {
        return req;
    }

    @Override
    public List<Boolean> EchoBoolList(List<Boolean> req) throws Exception {
        return req;
    }

    @Override
    public List<Byte> EchoByteList(List<Byte> req) throws Exception {
        return req;
    }

    @Override
    public List<Short> EchoInt16List(List<Short> req) throws Exception {
        return req;
    }

    @Override
    public List<Integer> EchoInt32List(List<Integer> req) throws Exception {
        return req;
    }

    @Override
    public List<Long> EchoInt64List(List<Long> req) throws Exception {
        return req;
    }

    @Override
    public List<Double> EchoDoubleList(List<Double> req) throws Exception {
        return req;
    }

    @Override
    public List<String> EchoStringList(List<String> req) throws Exception {
        return req;
    }

    @Override
    public List<byte[]> EchoBinaryList(List<byte[]> req) throws Exception {
        return req;
    }

    @Override
    public Map<Boolean, Boolean> EchoBool2BoolMap(Map<Boolean, Boolean> req) throws Exception {
        return req;
    }

    @Override
    public Map<Boolean, Byte> EchoBool2ByteMap(Map<Boolean, Byte> req) throws Exception {
        return req;
    }

    @Override
    public Map<Boolean, Short> EchoBool2Int16Map(Map<Boolean, Short> req) throws Exception {
        return req;
    }

    @Override
    public Map<Boolean, Integer> EchoBool2Int32Map(Map<Boolean, Integer> req) throws Exception {
        return req;
    }

    @Override
    public Map<Boolean, Long> EchoBool2Int64Map(Map<Boolean, Long> req) throws Exception {
        return req;
    }

    @Override
    public Map<Boolean, Double> EchoBool2DoubleMap(Map<Boolean, Double> req) throws Exception {
        return req;
    }

    @Override
    public Map<Boolean, String> EchoBool2StringMap(Map<Boolean, String> req) throws Exception {
        return req;
    }

    @Override
    public Map<Boolean, byte[]> EchoBool2BinaryMap(Map<Boolean, byte[]> req) throws Exception {
        return req;
    }

    @Override
    public EchoMultiBoolResponse EchoMultiBool(Boolean baseReq, List<Boolean> listReq, Map<Boolean, Boolean> mapReq) throws Exception {
        return new EchoMultiBoolResponse(baseReq, listReq, mapReq);
    }

    @Override
    public EchoMultiByteResponse EchoMultiByte(Byte baseReq, List<Byte> listReq, Map<Byte, Byte> mapReq) throws Exception {
        return new EchoMultiByteResponse(baseReq, listReq, mapReq);
    }

    @Override
    public EchoMultiInt16Response EchoMultiInt16(Short baseReq, List<Short> listReq, Map<Short, Short> mapReq) throws Exception {
        return new EchoMultiInt16Response(baseReq, listReq, mapReq);
    }

    @Override
    public EchoMultiInt32Response EchoMultiInt32(Integer baseReq, List<Integer> listReq, Map<Integer, Integer> mapReq) throws Exception {
        return new EchoMultiInt32Response(baseReq, listReq, mapReq);
    }

    @Override
    public EchoMultiInt64Response EchoMultiInt64(Long baseReq, List<Long> listReq, Map<Long, Long> mapReq) throws Exception {
        return new EchoMultiInt64Response(baseReq, listReq, mapReq);
    }

    @Override
    public EchoMultiDoubleResponse EchoMultiDouble(Double baseReq, List<Double> listReq, Map<Double, Double> mapReq) throws Exception {
        return new EchoMultiDoubleResponse(baseReq, listReq, mapReq);
    }

    @Override
    public EchoMultiStringResponse EchoMultiString(String baseReq, List<String> listReq, Map<String, String> mapReq) throws Exception {
        return new EchoMultiStringResponse(baseReq, listReq, mapReq);
    }
}