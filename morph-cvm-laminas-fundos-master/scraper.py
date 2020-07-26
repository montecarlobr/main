# -*- coding: utf-8 -*-

# This is a template for a Python scraper on morph.io (https://morph.io)
# including some code snippets below that you should find helpful

from __future__ import print_function
import scraperwiki
import lxml.html
import time


#
# # Read in a page

for code in range(20001, 40000):
    print(str(code), end=' ')
    time.sleep(.6)
    html = scraperwiki.scrape("http://cvmweb.cvm.gov.br/SWB/Sistemas/SCW/CPublica/CPublicaLamina.aspx?PK_PARTIC={}".format(str(code)))
    #
    # # Find something on the page using css selectors
    root = lxml.html.fromstring(html)

    pk_partic = code
    nome_admin = root.get_element_by_id("lblAdmin").text

    if nome_admin is None:
        continue    
    
    if "o foi encontrada nenhuma" in nome_admin:
        continue

    nome_fundo = root.get_element_by_id("lblNomeFundo").text
    cnpj_fundo = root.get_element_by_id("lblCnpj").text.replace('.', '').replace('-', '').replace('/', '')
    tipo_fundo = root.get_element_by_id("lblTipo").text
    codigo_cvm = root.get_element_by_id("lblCodCVM").text
    cnpj_admin = root.get_element_by_id("lblCnpjAdmin").text.replace('.', '').replace('-', '').replace('/', '')

    data = {
        "pk_partic": pk_partic,
        "nome_fundo": nome_fundo,
        "cnpj_fundo": cnpj_fundo,
        "tipo_fundo": tipo_fundo,
        "codigo_cvm": codigo_cvm,
        "nome_admin": nome_admin,
        "cnpj_admin": cnpj_admin
    }

    # Write out to the sqlite database using scraperwiki library
    scraperwiki.sqlite.save(unique_keys=['pk_partic'], data=data)
